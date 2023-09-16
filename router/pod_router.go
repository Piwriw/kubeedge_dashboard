package router

import (
	"github.com/gin-gonic/gin"
	"new-ec-dashboard/controller"
)

func PodRouter(r *gin.Engine) {
	podGroup := r.Group("/pod")
	{
		podGroup.GET("/pod/default/", controller.GetPodListHandler)
		podGroup.GET("/pod/:namespace/:name/", controller.GetPodHandler)
		podGroup.POST("/pod/", controller.CreateAPPHandler)
		podGroup.DELETE("/:namespace/:name/", controller.DeletePodHandler)
	}

}
