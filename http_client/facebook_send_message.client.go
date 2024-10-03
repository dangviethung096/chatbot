package http_client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dangviethung096/core"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/config"
)

type SendMessageToFacebookRequest struct {
	RecipientID string `json:"recipient_id"`
	Message     string `json:"message"`
}

type Recipient struct {
	ID string `json:"id"`
}

type Message struct {
	Text string `json:"text"`
}

type SendMessageToFacebookResponse struct {
	RecipientID string                     `json:"recipient_id"`
	MessageID   string                     `json:"message_id"`
	Error       SendMessageToFacebookError `json:"error"`
}

type SendMessageToFacebookError struct {
	Message   string `json:"message"`
	Type      string `json:"type"`
	Code      int64  `json:"code"`
	FBTraceID string `json:"fbtrace_id"`
}

func ResponseFacebookMessage(ctx core.Context, request SendMessageToFacebookRequest) core.Error {
	url := fmt.Sprintf("https://graph.facebook.com/v20.0/%d/messages", config.Value.Facebook.PageID)
	recipient := Recipient{
		ID: request.RecipientID,
	}

	recipientJson, err := json.Marshal(recipient)
	if err != nil {
		return core.NewError(http.StatusInternalServerError, err.Error())
	}

	message := Message{
		Text: request.Message,
	}

	messageJson, err := json.Marshal(message)
	if err != nil {
		return core.NewError(http.StatusInternalServerError, err.Error())
	}

	res := SendMessageToFacebookResponse{}

	_, coreError := Init(ctx).
		SetUrl(url).
		SetMethod(http.MethodPost).
		AddFormData("recipient", string(recipientJson)).
		AddFormData("message", string(messageJson)).
		AddFormData("messaging_type", "RESPONSE").
		AddFormData("access_token", config.Value.Facebook.Token).
		Request(&res)

	if coreError != nil {
		ctx.LogError("Error sending message to facebook: %s", coreError.Error())
		return coreError
	}

	if res.Error.Code != 0 {
		ctx.LogError("Error sending message to facebook: %v", res.Error)
		return core.NewError(http.StatusInternalServerError, res.Error.Message)
	}

	return nil
}
