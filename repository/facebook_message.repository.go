package repository

import (
	"github.com/dangviethung096/core"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/model"
)

func CreateFacebookMessage(ctx core.Context, message *model.FacebookMessage) (int64, core.Error) {
	err := core.SaveDataToDBWithoutPrimaryKey(ctx, message)
	if err != nil {
		ctx.LogError("Error creating facebook message: %v", err)
		return 0, err
	}
	return message.ID, nil
}

func GetFacebookMessageBySenderIDAndMessage(ctx core.Context, senderID string, message string) ([]model.FacebookMessage, core.Error) {
	messages, err := core.SelectListByFields(ctx, &model.FacebookMessage{}, map[string]interface{}{
		"sender":  senderID,
		"message": message,
	})
	if err != nil {
		ctx.LogError("Error getting facebook message: %v", err)
		return nil, err
	}

	if len(messages.([]model.FacebookMessage)) == 0 {
		ctx.LogError("Not found facebook message")
		return messages.([]model.FacebookMessage), core.ERROR_NOT_FOUND_IN_DB
	}

	return messages.([]model.FacebookMessage), nil
}
