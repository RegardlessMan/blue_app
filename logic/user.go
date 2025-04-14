package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"
)

// SignUp 用户注册
func SignUp(p *models.ParamSignUp) (err error) {

	// 检查用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	//插入用户
	user := &models.User{
		UserID:   snowflake.GenID(),
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.InsertUser(user)

}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}

	if err = mysql.Login(user); err != nil {
		return nil, err
	}
	token, err := jwt.GenerateToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
