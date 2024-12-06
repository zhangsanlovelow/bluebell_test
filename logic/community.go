package logic

import (
	"bullbell_test/dao/mysql"
	"bullbell_test/models"
)

// CommunityList 社区列表
func CommunityList() (communityList []*models.Community, err error) {
	return mysql.GetCommunityList()
}

// GetCommunityDetailByID 根据社区ID获取社区详情
func GetCommunityDetailByID(community_id int64) (communityDetail *models.CommunityDetail, err error) {
	return mysql.GetCommunityDetailByID(community_id)
}
