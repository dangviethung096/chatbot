package repository

import (
	"time"

	"github.com/dangviethung096/core"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/constant"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/model"
)

func GetZaloSessionByUserID(ctx core.Context, userID string) (*model.ZaloSession, core.Error) {
	var session model.ZaloSession
	err := core.SelectByField(ctx, &session, "sender_id", userID)
	if err != nil {
		ctx.LogError("GetZaloSessionByUserID error: %s", err.Error())
		return nil, err
	}

	return &session, nil
}

func InsertZaloMessage(ctx core.Context, message *model.ZaloMessage) (int64, core.Error) {
	message.CreatedAt = time.Now()
	err := core.SaveDataToDBWithoutPrimaryKey(ctx, message)
	if err != nil {
		ctx.LogError("InsertZaloMessage error: %s", err.Error())
		return constant.DEFAULT_INTEGER, err
	}

	return message.ID, nil
}

func InsertZaloSession(ctx core.Context, session *model.ZaloSession) (int64, core.Error) {
	session.CreatedAt = time.Now()
	session.UpdatedAt = time.Now()
	err := core.SaveDataToDBWithoutPrimaryKey(ctx, session)
	if err != nil {
		ctx.LogError("InsertZaloSession error: %s", err.Error())
		return constant.DEFAULT_INTEGER, err
	}

	return session.ID, nil
}

func UpdateZaloSession(ctx core.Context, session *model.ZaloSession) core.Error {
	session.UpdatedAt = time.Now()
	err := core.UpdateDataInDB(ctx, session)
	if err != nil {
		ctx.LogError("UpdateZaloSession error: %s", err.Error())
		return err
	}

	return nil
}
