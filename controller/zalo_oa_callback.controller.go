package controller

import (
	"net/http"

	"github.com/dangviethung096/core"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/config"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/errorpkg"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/http_client"
)

type ZaloOACallbackRequest struct {
}

type ZaloOACallbackResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func ZaloOACallback(ctx *core.HttpContext, request ZaloOACallbackRequest) (core.HttpResponse, core.HttpError) {
	ctx.LogInfo("Zalo OACallback request: %v", request)

	code := ctx.GetQueryParam("code")
	oaID := ctx.GetQueryParam("oa_id")
	state := ctx.GetQueryParam("state")
	codeChallenge := ctx.GetQueryParam("code_challenge")

	if state != config.Value.Zalo.State {
		ctx.LogError("State is invalid: %s", state)
		return nil, core.NewHttpError(http.StatusBadRequest, errorpkg.ERROR_CODE_BAD_REQUEST, "State is invalid", nil)
	}

	if codeChallenge != config.Value.Zalo.CodeChallenge {
		ctx.LogError("Code challenge is invalid: %s", codeChallenge)
		return nil, core.NewHttpError(http.StatusBadRequest, errorpkg.ERROR_CODE_BAD_REQUEST, "Code challenge is invalid", nil)
	}

	config.Value.Zalo.OaID = oaID
	config.Value.Zalo.OaCode = code
	config.WriteConfigFile()

	err := http_client.GetZaloToken(ctx, code)
	if err != nil {
		ctx.LogError("GetZaloToken error: %s", err.Error())
		return nil, core.NewHttpErrorFromError(err)
	}

	return core.NewDefaultHttpResponse(ZaloOACallbackResponse{
		AccessToken:  config.ZaloToken.AccessToken,
		ExpiresIn:    config.ZaloToken.ExpiresIn,
		RefreshToken: config.ZaloToken.RefreshToken,
	}), nil
}
