package http_client

import (
	"fmt"
	"net/http"

	"github.com/dangviethung096/core"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/config"
)

type CreateThreadResponse struct {
	ID string `json:"id"`
}

type openAICreateThreadResponse struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	CreatedAt int64  `json:"created_at"`
}

func CreateThread(ctx core.Context) (*CreateThreadResponse, core.Error) {
	res := openAICreateThreadResponse{}

	url := fmt.Sprintf("%s/v1/threads", config.Value.OpenAI.OpenAIUrl)
	_, err := Init(ctx).
		SetUrl(url).
		SetMethod(http.MethodPost).
		AddHeader(core.CONTENT_TYPE_KEY, core.JSON_CONTENT_TYPE).
		AddHeader("Authorization", fmt.Sprintf("Bearer %s", config.Value.OpenAI.ApiKey)).
		AddHeader("OpenAI-Beta", "assistants=v2").
		Request(&res)

	if err != nil {
		ctx.LogError("Request create thread fail: %v", err)
		return nil, err
	}

	ctx.LogInfo("Create new thread: %v", res)

	return &CreateThreadResponse{
		ID: res.ID,
	}, nil
}
