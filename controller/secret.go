package controller

import (
	"github.com/gin-gonic/gin"
	"new-ec-dashboard/service"
)

func GetSecretDetailOrList(c *gin.Context)  {
	namespace := c.Param("namespace")
	name := c.Query("name")
	keyword := c.Query("keyword")
	if name != "" {
		secretDetail, err := service.GetSecretDetail(namespace, name)
		if err != nil {
			ResponseErrorWithMsg(c, 200, "获取Secret详情失败")
			return
		}
		ResponseSuccess(c, secretDetail)
	} else {
		secretMapList, err := service.GetSecretList(namespace, keyword)
		if err != nil {
			ResponseErrorWithMsg(c, 200, "获取Secret列表详情失败")
			return
		}
		ResponseSuccess(c, secretMapList)
	}
}
