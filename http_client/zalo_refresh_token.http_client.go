package http_client

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/dangviethung096/core"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/config"
)

type RefreshZaloTokenResponse struct {
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token"`
	ExpiresIn             string `json:"expires_in"`
	RefreshTokenExpiresIn string `json:"refresh_token_expires_in"`
}

func RefreshZaloToken(ctx core.Context) core.Error {

	var res RefreshZaloTokenResponse

	_, err := Init(ctx).
		SetUrl("https://oauth.zaloapp.com/v4/oa/access_token").
		SetMethod(http.MethodPost).
		AddFormData("app_id", config.Value.Zalo.AppID).
		AddFormData("grant_type", "refresh_token").
		AddFormData("refresh_token", config.ZaloToken.RefreshToken).
		AddHeader("secret_key", config.Value.Zalo.AppSecret).
		Request(&res)

	if err != nil {
		ctx.LogError("RefreshZaloToken error = %s", err.Error())
		return err
	}

	ctx.LogInfo("RefreshZaloToken res = %+v", res)

	config.ZaloToken.AccessToken = res.AccessToken
	config.ZaloToken.RefreshToken = res.RefreshToken
	config.ZaloToken.ExpiresIn = res.ExpiresIn

	// Save to file
	f, originErr := os.Create(config.Value.Zalo.TokenFile)
	if originErr != nil {
		ctx.LogError("RefreshZaloToken error = %s", originErr.Error())
		return err
	}

	defer f.Close()

	originErr = json.NewEncoder(f).Encode(config.ZaloToken)
	if originErr != nil {
		ctx.LogError("RefreshZaloToken error = %s", originErr.Error())
		return err
	}

	return nil
}
