package http_client

import (
	"fmt"
	"net/http"

	"github.com/dangviethung096/core"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/config"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/constant"
)

type CreateMessageRequest struct {
	ThreadID string `json:"thread_id"`
	Content  string `json:"content"`
}

type CreateMessageResponse struct {
	ID string `json:"id"`
}

type openAICreateMessageResponseTextContent struct {
	Value string `json:"value"`
}

type openAICreateMessageResponseContent struct {
	Type string                                 `json:"type"`
	Text openAICreateMessageResponseTextContent `json:"text"`
}

type openAICreateMessageResponse struct {
	ID          string                               `json:"id"`
	Object      string                               `json:"object"`
	CreatedAt   int64                                `json:"created_at"`
	AssistantID *string                              `json:"assistant_id"`
	ThreadID    string                               `json:"thread_id"`
	RunID       *string                              `json:"run_id"`
	Role        string                               `json:"role"`
	Content     []openAICreateMessageResponseContent `json:"content"`
	Attachments []interface{}                        `json:"attachments"`
	Metadata    map[string]interface{}               `json:"metadata"`
}

type openAICreateMessageRequest struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func CreateMessage(ctx core.Context, request CreateMessageRequest) (*CreateMessageResponse, core.Error) {
	res := openAICreateMessageResponse{}
	url := fmt.Sprintf("%s/v1/threads/%s/messages", config.Value.OpenAI.OpenAIUrl, request.ThreadID)
	body := openAICreateMessageRequest{
		Role:    constant.OPENAI_ROLE_USER,
		Content: request.Content,
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
		ctx.LogError("Request create message fail: %v", err)
		return nil, err
	}

	ctx.LogInfo("Create new message: %v", res)

	return &CreateMessageResponse{
		ID: res.ID,
	}, nil
}
