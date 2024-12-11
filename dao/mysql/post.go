package mysql

import "bullbell_test/models"

// CreatePostHandler 创建帖子
func CreatePost(post *models.Post) (err error) {
	sqlStr := `insert into post(
		post_id,community_id,author_id,title,content,create_time) 
		values(?,?,?,?,?,?)
		`
	_, err = db.Exec(sqlStr, post.ID, post.CommunityID, post.AuthorID, post.Title, post.Content, post.CreateTime)

	return
}
