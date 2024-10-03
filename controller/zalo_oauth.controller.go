package controller

import (
	"net/http"

	"github.com/dangviethung096/core"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/config"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/errorpkg"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/http_client"
)

type ZaloOauthRequest struct {
}

type ZaloOauthResponse struct {
	Code string `json:"code"`
}

func ZaloOauth(ctx *core.HttpContext, request ZaloOauthRequest) (core.HttpResponse, core.HttpError) {
	//  https://yourdomain.com/abc?code=<AUTHORIZATION_CODE>&state=xxxx&code_challenge=xxxx
	code := ctx.GetQueryParam("code")
	state := ctx.GetQueryParam("state")

	if state != config.Value.Zalo.State {
		ctx.LogError("State is invalid: %s", state)
		return nil, core.NewHttpError(http.StatusBadRequest, errorpkg.ERROR_CODE_BAD_REQUEST, "State is invalid", nil)
	}

	err := http_client.GetZaloToken(ctx, code)
	if err != nil {
		ctx.LogError("GetZaloToken error: %s", err.Error())
		return nil, core.NewHttpErrorFromError(err)
	}

	return core.NewHttpResponse(200, ZaloOauthResponse{
		Code: code,
	}), nil
}
