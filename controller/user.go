package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLoginHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"access": "access123",
	})
	//p := new(models.ParamLogin)
	//if err := c.ShouldBindJSON(p); err != nil {
	//	zap.L().Error("Login with valid param", zap.Error(err))
	//	errs, ok := err.(validator.ValidationErrors)
	//	if !ok {
	//		ResponseError(c, CodeInvalidParam)
	//		return
	//	}
	//	ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
	//
	//	return
	//}
	//user, err := service.Login(p)
	//if err != nil {
	//	zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
	//	if errors.Is(err, mysql.ErrorUserNotExist) {
	//		ResponseError(c, CodeUserNotExist)
	//		return
	//	}
	//	ResponseError(c, CodeInvalidPassword)
	//	return
	//}
	//ResponseSuccess(c, gin.H{
	//	"user_id":   fmt.Sprintf("%d",user.UserID),
	//	"user_name": user.UserName,
	//	"token":     user.Token,
	//})
}
