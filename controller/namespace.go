package controller

import (
	"github.com/gin-gonic/gin"
	"new-ec-dashboard/service"
)

func GetAllNamespaces(c *gin.Context) {
	namespaces, err := service.GetAllNamespace()
	if err != nil {
		ResponseErrorWithMsg(c,500,"获取所有的namespace失败")
		return
	}
	ResponseSuccess(c,namespaces)
}
