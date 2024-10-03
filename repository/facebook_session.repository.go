package repository

import (
	"time"

	"github.com/dangviethung096/core"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/model"
)

func CreateFacebookSession(ctx core.Context, session *model.FacebookSession) (int64, core.Error) {
	now := time.Now()
	session.CreatedAt = now
	session.UpdatedAt = now
	err := core.SaveDataToDBWithoutPrimaryKey(ctx, session)
	if err != nil {
		ctx.LogError("Error creating facebook session: %v", err)
		return 0, err
	}
	return session.ID, nil
}

func GetSessionBySenderID(ctx core.Context, senderID string) (*model.FacebookSession, core.Error) {
	session := &model.FacebookSession{}
	err := core.SelectByField(ctx, session, "sender", senderID)
	if err != nil {
		ctx.LogError("Error getting facebook session: %v", err)
		return nil, err
	}
	return session, nil
}

func UpdateFacebookSession(ctx core.Context, session *model.FacebookSession) core.Error {
	session.UpdatedAt = time.Now()
	err := core.UpdateDataInDB(ctx, session)
	if err != nil {
		ctx.LogError("Error updating facebook session: %v", err)
		return err
	}
	return nil
}
