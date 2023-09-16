package router

import (
	"github.com/gin-gonic/gin"
	"new-ec-dashboard/controller"
)

func AIRouter(r *gin.Engine) {
	aiGroup := r.Group("/ai_model")
	{
	//	aiGroup.POST("/ai_model/", controller.CreateAIHandler)
		aiGroup.GET("/ai_model/", controller.GetListAIHandler)
		aiGroup.POST("/distribute/", controller.DistributeAiHandler)
	}

}
