package http_client

import (
	"net/http"

	"github.com/dangviethung096/core"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/config"
)

type GenerateContentRequest struct {
	Prompt string `json:"prompt"`
}

type GenerateContentResponse struct {
	Response string `json:"response"`
}

type googleGeminiGenerativeChatRequestPart struct {
	Text string `json:"text"`
}

type googleGeminiGenerativeChatRequestContent struct {
	Parts []googleGeminiGenerativeChatRequestPart `json:"parts"`
}

type googleGeminiGenerativeChatRequest struct {
	Contents []googleGeminiGenerativeChatRequestContent `json:"contents"`
}

type googleGeminiGenerativeChatResponseSafetyRating struct {
	Category    string `json:"category"`
	Probability string `json:"probability"`
}

type googleGeminiGenerativeChatResponsePart struct {
	Text string `json:"text"`
}

type googleGeminiGenerativeChatResponseContent struct {
	Parts []googleGeminiGenerativeChatResponsePart `json:"parts"`
	Role  string                                   `json:"role"`
}

type googleGeminiGenerativeChatResponseCandidate struct {
	Content       googleGeminiGenerativeChatResponseContent        `json:"content"`
	FinishReason  string                                           `json:"finishReason"`
	Index         int                                              `json:"index"`
	SafetyRatings []googleGeminiGenerativeChatResponseSafetyRating `json:"safetyRatings"`
}

type googleGeminiGenerativeChatResponseUsageMetadata struct {
	PromptTokenCount     int `json:"promptTokenCount"`
	CandidatesTokenCount int `json:"candidatesTokenCount"`
	TotalTokenCount      int `json:"totalTokenCount"`
}

type googleGeminiGenerativeChatResponse struct {
	Candidates    []googleGeminiGenerativeChatResponseCandidate   `json:"candidates"`
	UsageMetadata googleGeminiGenerativeChatResponseUsageMetadata `json:"usageMetadata"`
	Error         googleGeminiGenerativeChatResponseError         `json:"error"`
}

type googleGeminiGenerativeChatResponseFieldViolation struct {
	Description string `json:"description"`
}

type googleGeminiGenerativeChatResponseDetail struct {
	Type            string                                             `json:"@type"`
	FieldViolations []googleGeminiGenerativeChatResponseFieldViolation `json:"fieldViolations"`
}

type googleGeminiGenerativeChatResponseError struct {
	Code    int                                        `json:"code"`
	Message string                                     `json:"message"`
	Status  string                                     `json:"status"`
	Details []googleGeminiGenerativeChatResponseDetail `json:"details"`
}

func GenerateContent(ctx core.Context, request GenerateContentRequest) (GenerateContentResponse, core.Error) {
	requestBody := googleGeminiGenerativeChatRequest{
		Contents: []googleGeminiGenerativeChatRequestContent{
			{
				Parts: []googleGeminiGenerativeChatRequestPart{
					{Text: request.Prompt},
				},
			},
		},
	}

	res := googleGeminiGenerativeChatResponse{}
	_, err := Init(ctx).
		SetUrl("https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash-latest:generateContent").
		AddQuery("key", config.Value.Google.AiKey).
		SetMethod(http.MethodPost).
		SetBody(requestBody).
		Request(&res)

	if err != nil {
		ctx.LogError("GenerateContent error = %s", err.Error())
		return GenerateContentResponse{}, err
	}

	if res.Error.Code != 0 {
		ctx.LogError("GenerateContent error = %v", res.Error)
		return GenerateContentResponse{}, core.NewError(http.StatusInternalServerError, res.Error.Message)
	}

	message := res.Candidates[0].Content.Parts[0].Text
	ctx.LogInfo("GenerateContent message = %s", message)

	return GenerateContentResponse{Response: message}, nil
}
