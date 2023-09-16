package controller

import (
	"github.com/gin-gonic/gin"
	"new-ec-dashboard/service"
)

func GetCongFigMapDetailOrList(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Query("name")
	keyword := c.Query("keyword")
	if name != "" {
		cmDetail, err := service.GetConfigMapDetail(namespace, name)
		if err != nil {
			ResponseErrorWithMsg(c, 200, "获取ConfigMap详情失败")
			return
		}
		ResponseSuccess(c, cmDetail)
	} else {
		configMapList, err := service.GetConfigMapList(namespace, keyword)
		if err != nil {
			ResponseErrorWithMsg(c, 200, "获取ConfigMap列表详情失败")
			return
		}
		ResponseSuccess(c, configMapList)
	}
}
