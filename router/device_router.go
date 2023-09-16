package router

import (
	"github.com/gin-gonic/gin"
	"new-ec-dashboard/controller"
)

func DeviceRouter(r *gin.Engine) {
	deviceGroup := r.Group("/device")
	{
		deviceGroup.GET("/device/", controller.GetDeviceListHandler)
		deviceGroup.POST("/devicemodel/", controller.CreateDeviceModelHandler)
		deviceGroup.GET("/devicemodel/", controller.GetDeviceModelListHandler)
		deviceGroup.GET("/devicemodel/default/:name/", controller.GetDeviceModelHandler)
		deviceGroup.PUT("/devicemodel/", controller.UpdateDeviceModelHandler)
		deviceGroup.DELETE("/devicemodel/", controller.DeleteDeviceModelHandler)
		deviceGroup.POST("/device/", controller.CreateDeviceHandler)
		deviceGroup.GET("/device/default/:name/", controller.GetDeviceHandler)
		deviceGroup.DELETE("/device/", controller.DeleteDeviceHandler)
		deviceGroup.PUT("/device/",controller.UpdateDeviceHandler)
	}

}