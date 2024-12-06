package models

import "time"

// Community 社区
type Community struct {
	ID   int    `json:"community_id" db:"community_id"`
	Name string `json:"community_name" db:"community_name"`
}

// CommunityDetail 社区详情
type CommunityDetail struct {
	ID           int       `json:"community_id" db:"community_id"`
	Name         string    `json:"community_name" db:"community_name"`
	Introduction string    `json:"introduction,omitempty" db:"introduction"`
	CreateTime   time.Time `json:"create_time,omitempty" db:"create_time"`
	UpdateTime   time.Time `json:"update_time,omitempty" db:"update_time"`
}
