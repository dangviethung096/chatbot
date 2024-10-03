package http_client

import (
	"fmt"
	"net/http"

	"github.com/dangviethung096/core"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/config"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/constant"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/errorpkg"
)

type ResponseMessage struct {
	ID      string `json:"id"`
	Role    string `json:"role"`
	Content string `json:"content"`
}

type openaiListMessagesResponseTextContent struct {
	Value string `json:"value"`
}

type openaiListMessagesResponseContent struct {
	Type string                                `json:"type"`
	Text openaiListMessagesResponseTextContent `json:"text"`
}

type openaiListMessagesResponseMessage struct {
	ID          string                              `json:"id"`
	Object      string                              `json:"object"`
	CreatedAt   int64                               `json:"created_at"`
	AssistantID *string                             `json:"assistant_id"`
	ThreadID    string                              `json:"thread_id"`
	RunID       *string                             `json:"run_id"`
	Role        string                              `json:"role"`
	Content     []openaiListMessagesResponseContent `json:"content"`
	Attachments []interface{}                       `json:"attachments"`
	Metadata    map[string]interface{}              `json:"metadata"`
}

type openaiListMessagesResponse struct {
	Object  string                              `json:"object"`
	Data    []openaiListMessagesResponseMessage `json:"data"`
	FirstID string                              `json:"first_id"`
	LastID  string                              `json:"last_id"`
	HasMore bool                                `json:"has_more"`
}

func GetResponseMessage(ctx core.Context, threadID string) (*ResponseMessage, core.Error) {
	res := openaiListMessagesResponse{}

	url := fmt.Sprintf("%s/v1/threads/%s/messages", config.Value.OpenAI.OpenAIUrl, threadID)

	_, err := Init(ctx).
		SetUrl(url).
		SetMethod(http.MethodGet).
		AddHeader(core.CONTENT_TYPE_KEY, core.JSON_CONTENT_TYPE).
		AddHeader("Authorization", fmt.Sprintf("Bearer %s", config.Value.OpenAI.ApiKey)).
		AddHeader("OpenAI-Beta", "assistants=v2").
		Request(&res)

	if err != nil {
		ctx.LogError("ListMessages error: %v", err)
		return nil, err
	}

	if len(res.Data) == 0 || len(res.Data[0].Content) == 0 {
		ctx.LogError("Not found message in thread %s", threadID)
		return nil, errorpkg.ERROR_NOT_FOUND_RESPONSE_MESSAGE
	}

	response := ResponseMessage{
		ID:      res.Data[0].ID,
		Role:    res.Data[0].Role,
		Content: res.Data[0].Content[0].Text.Value,
	}

	if response.Role != constant.OPENAI_ROLE_ASSISTANT {
		ctx.LogError("Not found response message")
		return nil, errorpkg.ERROR_NOT_FOUND_RESPONSE_MESSAGE
	}

	return &response, nil
}
