package router

import (
	"github.com/gin-gonic/gin"
	"new-ec-dashboard/controller"
)

func IeRouter(r *gin.Engine) {
	ieGroup := r.Group("/ie_data")
	{
		ieGroup.POST("/ie_data", controller.PostIeDataHandler)
		ieGroup.GET("/ie_data", controller.GetIeDataListHandler)
	}

}
