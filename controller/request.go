package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)
const CtxUserIDKey="userID"
var ErrorNotLogin=errors.New("用户未登录")

// getCurrentUser : 获取当前登录用户的ID
func getCurrentUser(c *gin.Context)(userID int64,err error)  {
	uid,ok:=c.Get(CtxUserIDKey)
	if !ok {
		err=ErrorNotLogin
		return
	}
	userID,ok=uid.(int64)
	if!ok{
		err=ErrorNotLogin
		return
	}
	return
}
// getPageInfo: 获取分页
func getPageInfo(c *gin.Context)(int64,int64){
	// 获取分页参数
	pageNumberStr:=c.Query("page")
	pageSizeStr:=c.Query("size")
	var (
		page int64
		size int64
		err error
	)
	page,err=strconv.ParseInt(pageNumberStr,10,64)
	if err != nil {
		page=1
	}
	size,err=strconv.ParseInt(pageSizeStr,10,64)
	if err != nil {
		size=10
	}
	return page,size
}