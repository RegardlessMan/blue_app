package mysql

import (
	"database/sql"
	"errors"
	"go.uber.org/zap"
	"web_app/models"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := `select community_id, community_name from community`
	if err := db.Select(&communityList, sqlStr); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

func GetCommunityDetailByID(id int64) (communityDetail *models.CommunityDetail, err error) {
	sqlStr := `select community_id,community_name,introduction,create_time from community where community_id = ?`
	communityDetail = new(models.CommunityDetail)
	if err := db.Get(communityDetail, sqlStr, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrorInvalidId
		}
	}
	return communityDetail, nil
}
