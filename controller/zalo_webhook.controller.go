package controller

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/dangviethung096/core"
	"github.com/russross/blackfriday/v2"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/config"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/constant"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/errorpkg"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/http_client"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/model"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/repository"
)

type ZaloWebhookRequest struct {
	AppID  string `json:"app_id"`
	Sender struct {
		ID string `json:"id"`
	} `json:"sender"`
	UserIDByApp string `json:"user_id_by_app"`
	Recipient   struct {
		ID string `json:"id"`
	} `json:"recipient"`
	EventName string `json:"event_name"`
	Message   struct {
		ConversationID string `json:"conversation_id"`
		Text           string `json:"text"`
		MsgID          string `json:"msg_id"`
	} `json:"message"`
	Timestamp string `json:"timestamp"`
}

func ZaloWebhook(ctx *core.HttpContext, request ZaloWebhookRequest) (core.HttpResponse, core.HttpError) {
	ctx.LogInfo("Zalo webhook received: %#v", request)

	ctx.EndResponse(200, nil, nil)

	ctx.LogInfo("Receive message from user: %s", request.Message.Text)

	if request.EventName != constant.ZALO_EVENT_ANONYMOUS_SEND_TEXT && request.EventName != constant.ZALO_EVENT_USER_SEND_TEXT {
		ctx.LogError("Invalid event name: %s", request.EventName)
		return nil, nil
	}

	// Check if message start with '#'
	if !strings.HasPrefix(request.Message.Text, "#") {
		ctx.LogInfo("Message not start with #: do not call the chatbot")
		return nil, nil
	}

	request.Message.Text = strings.TrimPrefix(request.Message.Text, "#")
	if request.Message.Text == "hello" || request.Message.Text == "hi" || request.Message.Text == "alo" {
		request.Message.Text = "Xin chào"
	}

	// Get zalo session
	zaloSession, err := repository.GetZaloSessionByUserID(ctx, request.Sender.ID)
	if err != nil {
		if err != core.ERROR_NOT_FOUND_IN_DB {
			ctx.LogError("GetZaloSessionByUserID error: sender_id = %s, error = %v", request.Sender.ID, err)
			return nil, nil
		}

		// Create a new thread
		threadInfo, err := http_client.CreateThread(ctx)
		if err != nil {
			ctx.LogError("CreateThread error: %s", err.Error())
			return nil, nil
		}

		// Create new zalo session
		zaloSession = &model.ZaloSession{
			SenderID: request.Sender.ID,
			ThreadID: threadInfo.ID,
		}

		_, err = repository.InsertZaloSession(ctx, zaloSession)
		if err != nil {
			ctx.LogError("CreateZaloSession error: %s", err.Error())
			return nil, nil
		}
	}

	// Check if thread is expired
	if zaloSession.UpdatedAt.Before(time.Now().Add(-constant.OPENAI_EXPIRED_TIME_OF_THREAD)) {
		// Create a new thread
		ctx.LogInfo("Thread is expired, create a new thread for sender: %s", request.Sender.ID)
		threadInfo, err := http_client.CreateThread(ctx)
		if err != nil {
			ctx.LogError("CreateThread error: %s", err.Error())
			return nil, nil
		}

		zaloSession.ThreadID = threadInfo.ID
		err = repository.UpdateZaloSession(ctx, zaloSession)
		if err != nil {
			ctx.LogError("CreateZaloSession error: %s", err.Error())
			return nil, nil
		}
	}

	// Add message to thread
	_, err = http_client.CreateMessage(ctx, http_client.CreateMessageRequest{
		ThreadID: zaloSession.ThreadID,
		Content:  request.Message.Text,
	})
	if err != nil {
		ctx.LogError("CreateMessage error: %s", err.Error())
		return nil, nil
	}

	// Run assistant
	_, err = http_client.RunAssistant(ctx, http_client.RunAssistantRequest{
		ThreadID:    zaloSession.ThreadID,
		AssistantID: config.Value.OpenAI.AssistantID,
	})
	if err != nil {
		ctx.LogError("RunAssistant error: %s", err.Error())
		return nil, nil
	}

	// List message in thread
	var responseMessage *http_client.ResponseMessage
	for i := 0; i < 3; i++ {
		responseMessage, err = http_client.GetResponseMessage(ctx, zaloSession.ThreadID)
		if err != nil {
			if err == errorpkg.ERROR_NOT_FOUND_RESPONSE_MESSAGE {
				time.Sleep(time.Second * 3)
				continue
			}
			ctx.LogError("ListMessages error: %s", err.Error())
			return nil, nil
		}

		break
	}

	resMessage := string(blackfriday.Run([]byte(responseMessage.Content), blackfriday.WithNoExtensions()))
	resMessage = htmlToPlainText(resMessage)
	resMessage = strings.ReplaceAll(resMessage, "\n\n", "\n")
	// remove all text in 【any text】
	resMessage = removeTextInBracket(resMessage)
	resMessage += "\n\nNội dung được tạo ra từ MobifoneKv4 AI Chatbot"

	messageRequest := http_client.SendZaloMessageRequest{
		Recipient: http_client.SendZaloMessageRequestRecipient{
			ConversationID: request.Message.ConversationID,
		},
		Message: http_client.SendZaloMessageRequestMessage{
			Text: resMessage,
		},
	}

	if request.EventName == constant.ZALO_EVENT_ANONYMOUS_SEND_TEXT {
		messageRequest.Recipient.AnonymousID = request.Sender.ID
		err = http_client.SendZaloMessageToAnonymous(ctx, messageRequest)
	} else {
		messageRequest.Recipient.UserID = request.Sender.ID
		err = http_client.SendZaloMessage(ctx, messageRequest)
	}

	if err != nil {
		ctx.LogError("SendZaloMessage error: %s", err.Error())
		return nil, nil
	}

	zaloMessage := model.ZaloMessage{
		SenderID:    request.Recipient.ID,
		Message:     resMessage,
		RecipientID: request.Sender.ID,
		ThreadID:    zaloSession.ThreadID,
		MessageID:   responseMessage.ID,
	}

	_, err = repository.InsertZaloMessage(ctx, &zaloMessage)
	if err != nil {
		ctx.LogError("InsertZaloMessage error: %s", err.Error())
		return nil, nil
	}

	ctx.LogInfo("Chatbot response message to user: user_id = %s, message = %s", request.Sender.ID, resMessage)

	return nil, nil
}

func removeTextInBracket(text string) string {
	newMessage := ""
	ignore := false
	for _, char := range text {
		if char == '【' {
			ignore = true
		}

		if char == '】' {
			ignore = false
			continue
		}

		if ignore {
			continue
		}
		newMessage += string(char)
	}

	return newMessage
}

func htmlToPlainText(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return html // fallback to original HTML if parsing fails
	}
	return doc.Text()
}
