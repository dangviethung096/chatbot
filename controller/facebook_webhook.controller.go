package controller

import (
	"net/http"
	"strings"
	"time"

	"github.com/dangviethung096/core"
	"github.com/russross/blackfriday/v2"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/config"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/constant"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/errorpkg"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/http_client"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/model"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/repository"
)

type FacebookWebhookRequest struct {
	Object string                        `json:"object"`
	Entry  []FacebookWebhookRequestEntry `json:"entry"`
}

type FacebookWebhookRequestEntry struct {
	Time      int64                             `json:"time"`
	ID        string                            `json:"id"`
	Messaging []FacebookWebhookRequestMessaging `json:"messaging"`
}

type FacebookWebhookRequestMessaging struct {
	Sender    FacebookWebhookRequestUser    `json:"sender"`
	Recipient FacebookWebhookRequestUser    `json:"recipient"`
	Timestamp int64                         `json:"timestamp"`
	Message   FacebookWebhookRequestMessage `json:"message"`
}

type FacebookWebhookRequestUser struct {
	ID string `json:"id"`
}

type FacebookWebhookRequestMessage struct {
	MID  string `json:"mid"`
	Text string `json:"text"`
}

type FacebookWebhookResponse struct {
}

const (
	DEFAULT_DUPLICATE_MESSAGE_TIME = -50 * time.Minute
	DEFAULT_TIME_TO_RUN_ASSISTANT  = -3 * time.Minute
)

func FacebookWebhook(ctx *core.HttpContext, request FacebookWebhookRequest) (core.HttpResponse, core.HttpError) {
	ctx.LogInfo("Facebook received request: %#v", request)
	ctx.EndResponse(http.StatusOK, nil, []byte("EVENT_RECEIVED"))

	for _, entry := range request.Entry {
		for _, messaging := range entry.Messaging {
			message := messaging.Message.Text
			senderID := messaging.Sender.ID
			recipientID := messaging.Recipient.ID
			ctx.LogInfo("Received message: %s, from %s to %s", message, senderID, recipientID)
			if message == constant.BLANK {
				continue
			}

			if message == "hello" || message == "hi" || message == "hey" || message == "alo" {
				message = "Xin chào"
			}

			entryTime := time.UnixMilli(entry.Time)
			now := time.Now()
			if entryTime.Before(now.Add(DEFAULT_TIME_TO_RUN_ASSISTANT)) {
				ctx.LogInfo("Message is too old")
				continue
			}

			handleMessage(ctx, message, senderID, recipientID)
		}
	}

	return nil, nil
}

func handleMessage(ctx *core.HttpContext, message string, senderID string, recipientID string) {
	ctx.LogInfo("Handling message: %s, from %s to %s", message, senderID, recipientID)
	fbMessages, _ := repository.GetFacebookMessageBySenderIDAndMessage(ctx, senderID, message)
	if len(fbMessages) > 0 {
		for _, fbMessage := range fbMessages {
			now := time.Now()
			if fbMessage.CreatedAt.Before(now.Add(DEFAULT_DUPLICATE_MESSAGE_TIME)) {
				ctx.LogInfo("Message is too new: %#v", fbMessage)
				return
			}
		}
	}

	session, err := repository.GetSessionBySenderID(ctx, senderID)
	if err != nil {
		if err != core.ERROR_NOT_FOUND_IN_DB {
			ctx.LogError("Error getting facebook session: %v", err)
			return
		}

		thread, err := http_client.CreateThread(ctx)
		if err != nil {
			ctx.LogError("Error creating thread: %v", err)
			return
		}

		session = &model.FacebookSession{
			Sender:    senderID,
			ThreadID:  thread.ID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		_, err = repository.CreateFacebookSession(ctx, session)
		if err != nil {
			ctx.LogError("Error creating facebook session: %v", err)
			return
		}
	}

	// save message to database
	fbMessage := model.FacebookMessage{
		Sender:    senderID,
		Receiver:  recipientID,
		Message:   message,
		ThreadID:  session.ThreadID,
		CreatedAt: time.Now(),
	}

	_, err = repository.CreateFacebookMessage(ctx, &fbMessage)
	if err != nil {
		ctx.LogError("Error creating facebook message: %v", err)
	}

	thread, err := http_client.CreateThread(ctx)
	if err != nil {
		ctx.LogError("Error creating thread: %v", err)
		return
	}

	session.ThreadID = thread.ID
	err = repository.UpdateFacebookSession(ctx, session)
	if err != nil {
		ctx.LogError("Error updating facebook session: %v", err)
		return
	}

	// Add message to thread
	_, err = http_client.CreateMessage(ctx, http_client.CreateMessageRequest{
		ThreadID: session.ThreadID,
		Content:  message,
	})
	if err != nil {
		ctx.LogError("CreateMessage error: %s", err.Error())
		return
	}

	// Run assistant
	_, err = http_client.RunAssistant(ctx, http_client.RunAssistantRequest{
		ThreadID:    session.ThreadID,
		AssistantID: config.Value.OpenAI.AssistantID,
	})
	if err != nil {
		ctx.LogError("RunAssistant error: %s", err.Error())
		return
	}

	// List message in thread
	var responseMessage *http_client.ResponseMessage
	for i := 0; i < 3; i++ {
		responseMessage, err = http_client.GetResponseMessage(ctx, session.ThreadID)
		if err != nil {
			if err == errorpkg.ERROR_NOT_FOUND_RESPONSE_MESSAGE {
				time.Sleep(time.Second * 3)
				continue
			}
			ctx.LogError("ListMessages error: %s", err.Error())
			return
		}

		break
	}

	resMessage := string(blackfriday.Run([]byte(responseMessage.Content), blackfriday.WithNoExtensions()))
	resMessage = htmlToPlainText(resMessage)
	resMessage = strings.ReplaceAll(resMessage, "\n\n", "\n")
	// remove all text in 【any text】
	resMessage = removeTextInBracket(resMessage)
	resMessage += "\n\nNội dung được tạo ra từ HPW AI Chatbot"

	err = http_client.ResponseFacebookMessage(ctx, http_client.SendMessageToFacebookRequest{
		RecipientID: senderID,
		Message:     resMessage,
	})
	if err != nil {
		ctx.LogError("Error response message to facebook: %s", err.Error())
		return
	}

	fbMessage = model.FacebookMessage{
		Sender:    recipientID,
		Receiver:  senderID,
		Message:   resMessage,
		ThreadID:  session.ThreadID,
		CreatedAt: time.Now(),
	}

	_, err = repository.CreateFacebookMessage(ctx, &fbMessage)
	if err != nil {
		ctx.LogError("Error creating facebook message: %v", err)
	}
}
