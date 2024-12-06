package mysql

import (
	"bullbell_test/models"
	"database/sql"

	"go.uber.org/zap"
)

// CommunityList 社区列表
func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := `select community_id,community_name from community`

	err = db.Select(&communityList, sqlStr)
	if err != sql.ErrNoRows {
		zap.L().Error("GetCommunityList failed", zap.Error(err))
	}
	return
}

// GetCommunityDetailByID 根据社区ID获取社区详情
func GetCommunityDetailByID(community_id int64) (communityDetail *models.CommunityDetail, err error) {
	communityDetail = new(models.CommunityDetail)
	sqlStr := `select community_id,community_name,introduction,create_time from community where community_id=?`
	if err = db.Get(communityDetail, sqlStr, community_id); err != nil {
		zap.L().Error("GetCommunityDetailByID failed", zap.Error(err))
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return communityDetail, err
}
