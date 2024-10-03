package controller

import (
	"crypto/sha256"
	"encoding/base64"

	"github.com/dangviethung096/core"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/config"
)

type ZaloCodeChallengeRequest struct {
	CodeVerifier string `json:"code_verifier"`
}

type ZaloCodeChallengeResponse struct {
	CodeChallenge string `json:"code_challenge"`
}

func ZaloCodeChallenge(ctx *core.HttpContext, request ZaloCodeChallengeRequest) (core.HttpResponse, core.HttpError) {
	hash := sha256.New()
	hash.Write([]byte(request.CodeVerifier))
	sha256Hash := hash.Sum(nil)

	// Encode the hash using Base64
	codeChallenge := base64.RawURLEncoding.EncodeToString(sha256Hash)

	// Create response
	response := ZaloCodeChallengeResponse{
		CodeChallenge: codeChallenge,
	}

	config.Value.Zalo.CodeChallenge = codeChallenge
	config.Value.Zalo.CodeVerifier = request.CodeVerifier
	config.WriteConfigFile()

	return core.NewDefaultHttpResponse(response), nil
}
