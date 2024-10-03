package http_client

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/dangviethung096/core"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/config"
)

type zaloResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        string `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	ErrorName        string `json:"error_name"`
	ErrorReason      string `json:"error_reason"`
	RefDoc           string `json:"ref_doc"`
	ErrorDescription string `json:"error_description"`
	Error            int64  `json:"error"`
}

func GetZaloToken(ctx core.Context, code string) core.Error {
	ctx.LogInfo("GetZaloClient code = %s", code)

	var res zaloResponse
	_, err := Init(ctx).
		SetUrl("https://oauth.zaloapp.com/v4/oa/access_token").
		SetMethod(http.MethodPost).
		AddFormData("app_id", config.Value.Zalo.AppID).
		AddFormData("grant_type", "authorization_code").
		AddFormData("code", code).
		AddFormData("code_verifier", config.Value.Zalo.CodeVerifier).
		AddHeader("secret_key", config.Value.Zalo.AppSecret).
		Request(&res)

	if err != nil {
		ctx.LogError("GetZaloClient error = %s", err.Error())
		return err
	}

	if res.Error != 0 {
		ctx.LogError("GetZaloClient error = %v", res)
		return core.NewError(http.StatusInternalServerError, res.ErrorName)
	}

	ctx.LogInfo("GetZaloClient res = %+v", res)
	// Assign zalo token
	config.ZaloToken.AccessToken = res.AccessToken
	config.ZaloToken.RefreshToken = res.RefreshToken
	config.ZaloToken.ExpiresIn = res.ExpiresIn

	// Save to file
	f, originErr := os.Create(config.Value.Zalo.TokenFile)
	if originErr != nil {
		ctx.LogError("GetZaloClient error = %s", originErr.Error())
		return core.NewError(http.StatusInternalServerError, originErr.Error())
	}

	defer f.Close()

	originErr = json.NewEncoder(f).Encode(config.ZaloToken)
	if originErr != nil {
		ctx.LogError("GetZaloClient error = %s", originErr.Error())
		return core.NewError(http.StatusInternalServerError, originErr.Error())
	}

	ctx.LogInfo("GetZaloClient success")
	return nil
}
