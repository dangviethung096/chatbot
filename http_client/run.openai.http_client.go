package http_client

import (
	"fmt"
	"net/http"

	"github.com/dangviethung096/core"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/config"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/constant"
)

type RunAssistantRequest struct {
	ThreadID    string
	AssistantID string
}

type RunAssistantResponse struct {
	ID string
}

type openAIRunAssistantResponse struct {
	ID          string `json:"id"`
	Object      string `json:"object"`
	CreatedAt   int64  `json:"created_at"`
	AssistantID string `json:"assistant_id"`
	ThreadID    string `json:"thread_id"`
	Status      string `json:"status"`
}

type openAIRunAssistantRequest struct {
	AssistantID        string                                      `json:"assistant_id"`
	TruncationStrategy openAIRunAssistantRequestTruncationStrategy `json:"truncation_strategy"`
}

type openAIRunAssistantRequestTruncationStrategy struct {
	Type         string `json:"type"`
	LastMessages int64  `json:"last_messages"`
}

func RunAssistant(ctx core.Context, request RunAssistantRequest) (*RunAssistantResponse, core.Error) {
	res := openAIRunAssistantResponse{}

	url := fmt.Sprintf("%s/v1/threads/%s/runs", config.Value.OpenAI.OpenAIUrl, request.ThreadID)
	body := openAIRunAssistantRequest{
		AssistantID: config.Value.OpenAI.AssistantID,
		TruncationStrategy: openAIRunAssistantRequestTruncationStrategy{
			Type:         "last_messages",
			LastMessages: constant.OPENAI_NUMBER_OF_MESSAGE_IN_THREAD,
		},
	}

	_, err := Init(ctx).
		SetUrl(url).
		SetMethod(http.MethodPost).
		AddHeader(core.CONTENT_TYPE_KEY, core.JSON_CONTENT_TYPE).
		AddHeader("Authorization", fmt.Sprintf("Bearer %s", config.Value.OpenAI.ApiKey)).
		AddHeader("OpenAI-Beta", "assistants=v2").
		SetBody(body).
		Request(&res)

	if err != nil {
		ctx.LogError("Request run assistant fail: %v", err)
		return nil, err
	}

	ctx.LogInfo("Run assistant success: %v", res)

	return &RunAssistantResponse{
		ID: res.ID,
	}, nil
}
