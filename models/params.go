package models

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type VoteData struct {
	// UserID 从请求中获取当前的用户
	PostID    int64 `json:"post_id,string"`   // 贴子id
	Direction int   `json:"direction,string"` // 赞成票(1)还是反对票(-1)
}
