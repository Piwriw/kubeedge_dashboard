package router

import (
	"github.com/gin-gonic/gin"
	"new-ec-dashboard/controller"
)

func NodeRouter(r *gin.Engine) {
	nodeGroup := r.Group("/node")
	{
		nodeGroup.GET("/node/:nodename/", controller.GetNodeByNameHandler)
		nodeGroup.GET("/node/", controller.GetNodeListHandler)
		nodeGroup.GET("/metrics/:nodename/", controller.GetNodeMetricsHandler)
		nodeGroup.GET("/join/", controller.GetJoinTokenHandler)
		nodeGroup.PUT("/label/", controller.CreateNodeLabelHandler)
		nodeGroup.DELETE("/label/", controller.DeleteNodeLabelHandler)
	}

}
