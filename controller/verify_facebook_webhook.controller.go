package controller

import (
	"net/http"

	"github.com/dangviethung096/core"
)

type VerifyFacebookWebhookRequest struct {
}

type VerifyFacebookWebhookResponse struct {
}

func VerifyFacebookWebhook(ctx *core.HttpContext, request VerifyFacebookWebhookRequest) (core.HttpResponse, core.HttpError) {
	ctx.LogInfo("Facebook received request: %#v", request)

	mode := ctx.GetQueryParam("hub.mode")
	challenge := ctx.GetQueryParam("hub.challenge")
	token := ctx.GetQueryParam("hub.verify_token")

	ctx.LogInfo("Facebook webhook verify: mode=%s, challenge=%s, token=%s", mode, challenge, token)

	if mode == "subscribe" && token == "chatbot" {
		ctx.LogInfo("Facebook webhook verified")
		ctx.EndResponse(http.StatusOK, nil, []byte(challenge))
	} else {
		ctx.LogInfo("Facebook webhook not verified: mode=%s, token=%s", mode, token)
		ctx.EndResponse(http.StatusForbidden, nil, nil)
	}

	return nil, nil
}
