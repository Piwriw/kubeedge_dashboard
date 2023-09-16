package router

import (
	"github.com/gin-gonic/gin"
	"new-ec-dashboard/controller"
)

func UserRouter(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", controller.UserLoginHandler)
	}

}
