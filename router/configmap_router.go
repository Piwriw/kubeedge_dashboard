package router

import (
	"github.com/gin-gonic/gin"
	"new-ec-dashboard/controller"
)

func ConfigMapRouter(r *gin.Engine) {
	aiGroup := r.Group("/configmap")
	{
		aiGroup.GET("/:namespace", controller.GetCongFigMapDetailOrList)
	}

}