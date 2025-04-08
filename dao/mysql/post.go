package mysql

import "web_app/models"

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(post_id, title, content, author_id, community_id) values (?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return err
}

func GetPostDetail(pid int64) (p *models.Post, err error) {
	p = &models.Post{}
	sqlStr := `select post_id, title, content, author_id, community_id from post where post_id=?`
	err = db.Get(p, sqlStr, pid)
	return
}

func GetPostList(page int64, size int64) (posts []*models.Post, err error) {
	posts = make([]*models.Post, 0, size)
	sqlStr := `select post_id, title, content, author_id, community_id from post limit ?,?`
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}
