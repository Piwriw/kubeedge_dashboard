package router

import (
	"github.com/gin-gonic/gin"
	"new-ec-dashboard/controller"
)

func SecretRouter(r *gin.Engine) {
	aiGroup := r.Group("/secret")
	{
		aiGroup.GET("/:namespace", controller.GetSecretDetailOrList)
	}

}
