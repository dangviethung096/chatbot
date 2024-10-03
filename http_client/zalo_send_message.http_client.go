package http_client

import (
	"net/http"

	"github.com/dangviethung096/core"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/config"
)

type SendZaloMessageRequest struct {
	Recipient SendZaloMessageRequestRecipient `json:"recipient"`
	Message   SendZaloMessageRequestMessage   `json:"message"`
}

type SendZaloMessageRequestRecipient struct {
	UserID         string `json:"user_id"`
	ConversationID string `json:"conversation_id,omitempty"`
	AnonymousID    string `json:"anonymous_id,omitempty"`
}

type SendZaloMessageRequestMessage struct {
	Text string `json:"text"`
}

type SendZaloMessageResponse struct {
	Data struct {
		Quota struct {
			QuotaType   string `json:"quota_type"`
			Remain      int    `json:"remain"`
			Total       int    `json:"total"`
			ExpiredDate string `json:"expired_date"`
			OwnerType   string `json:"owner_type"`
			OwnerID     string `json:"owner_id"`
		} `json:"quota"`
		MessageID string `json:"message_id"`
		UserID    string `json:"user_id"`
	} `json:"data"`
	Error   int    `json:"error"`
	Message string `json:"message"`
}

func SendZaloMessage(ctx core.Context, request SendZaloMessageRequest) core.Error {
	var res SendZaloMessageResponse

	_, err := Init(ctx).
		SetUrl("https://openapi.zalo.me/v3.0/oa/message/cs").
		SetMethod(http.MethodPost).
		AddHeader("access_token", config.ZaloToken.AccessToken).
		AddHeader("Content-Type", core.JSON_CONTENT_TYPE).
		SetBody(request).
		Request(&res)

	if err != nil {
		ctx.LogError("GetZaloClient error = %s", err.Error())
		return err
	}

	ctx.LogInfo("SendZaloMessage res = %#v", res)
	if res.Error != 0 {
		ctx.LogError("SendZaloMessage error = %#v", res)
		return core.NewError(http.StatusInternalServerError, res.Message)
	}

	return nil
}

func SendZaloMessageToAnonymous(ctx core.Context, request SendZaloMessageRequest) core.Error {
	var res SendZaloMessageResponse

	_, err := Init(ctx).
		SetUrl("https://openapi.zalo.me/v2.0/oa/message").
		SetMethod(http.MethodPost).
		AddHeader("access_token", config.ZaloToken.AccessToken).
		AddHeader("Content-Type", core.JSON_CONTENT_TYPE).
		SetBody(request).
		Request(&res)

	if err != nil {
		ctx.LogError("GetZaloClient error = %s", err.Error())
		return err
	}

	ctx.LogInfo("SendZaloMessage res = %#v", res)
	if res.Error != 0 {
		ctx.LogError("SendZaloMessage error = %#v", res)
		return core.NewError(http.StatusInternalServerError, res.Message)
	}

	return nil
}
