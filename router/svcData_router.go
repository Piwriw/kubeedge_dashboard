package router

import (
	"github.com/gin-gonic/gin"
	"new-ec-dashboard/controller"
)

func SvcDataRouter(r *gin.Engine) {
	svcDataGroup := r.Group("/svc_data")
	{
		svcDataGroup.POST("/svc/", controller.PostSvcDataListHandler)
		svcDataGroup.GET("/svc_data/", controller.GetSvcDataListHandler)

		//svcDataGroup.GET("/svc_data/:id",controller.)
	}

}
