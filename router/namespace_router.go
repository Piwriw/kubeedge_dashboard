package router

import (
	"github.com/gin-gonic/gin"
	"new-ec-dashboard/controller"
)

func NSRouter(r *gin.Engine) {
	ieGroup := r.Group("/namespace")
	{
		ieGroup.GET("/", controller.GetAllNamespaces)
	}

}


