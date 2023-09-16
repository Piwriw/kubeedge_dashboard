package service

import (
	"new-ec-dashboard/dao/mysql"
	"new-ec-dashboard/models"
	"new-ec-dashboard/pkg/jwt"
)

// Login : 登录
func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		UserName: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	// 生成JWT
	token, err := jwt.GenToken(user.UserID, user.UserName)
	if err != nil {
		return
	}
	user.Token = token
	return
}
