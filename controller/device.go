package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"new-ec-dashboard/models"
	"new-ec-dashboard/service"
)

// GetDeviceListHandler : 获取Device详情
// @Summary Get Device接口
// @Tags Device相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /device/device/ [post]
func GetDeviceListHandler(c *gin.Context) {
	deviceList, err := service.GetDeviceList()
	if err != nil {
		zap.L().Error("service.GetDeviceList() is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, deviceList)
}

// CreateDeviceModelHandler : 获取Device Model详情
// @Summary Get Device Model详情接口
// @Tags Device Model详情相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /device/devicemodel/ [post]
func CreateDeviceModelHandler(c *gin.Context) {
	deviceModel := new(models.DeviceModelParam)
	if err := c.ShouldBindJSON(&deviceModel); err != nil {
		zap.L().Error("CreateDeviceModelHandler post invalid param!")
		ResponseError(c, CodeInvalidParam)
		return
	}
	model, err := service.CreateDeviceModel(deviceModel)
	if err != nil {
		zap.L().Error("service.CreateDeviceModel(deviceModel) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, model)
}

// GetDeviceModelListHandler : 获取Device Model列表详情
// @Summary Get Device Model详情接口
// @Tags Device Model详情相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /device/devicemodel/ [get]
func GetDeviceModelListHandler(c *gin.Context) {
	deviceModelList, err := service.GetDeviceModelList()
	if err != nil {
		zap.L().Error("service.GetDeviceModelList() is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, deviceModelList)
}

// GetDeviceModelHandler : 获取Device Model详情
// @Summary Get Device Model详情接口
// @Tags Device Model详情相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /device/devicemodel/default/:name/ [get]
func GetDeviceModelHandler(c *gin.Context) {
	name := c.Param("name")
	deviceModel, err := service.GetDeviceModel(name)
	if err != nil {
		zap.L().Error("service.GetDeviceModel(name) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, deviceModel)
}

// UpdateDeviceModelHandler : 更新Device Model详情
// @Summary Update Device Model详情接口
// @Tags Device Model 相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /device/devicemodel/ [put]
func UpdateDeviceModelHandler(c *gin.Context) {
	p := new(models.DeviceModelParam)
	err := c.ShouldBindJSON(p)
	if err != nil {
		zap.L().Error("UpdateDeviceModelHandler c.ShouldBindJSON(p) is failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	_, err = service.UpdateDeviceModel(p)
	if err != nil {
		zap.L().Error("service.UpdateDeviceModel(p) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}

// DeleteDeviceModelHandler : 删除Device Model详情
// @Summary DELETE Device Model详情接口
// @Tags Device Model 相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /device/devicemodel/ [DELETE]
func DeleteDeviceModelHandler(c *gin.Context) {
	p := new(models.DeviceModelParam)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("DeleteDeviceModelHandler with valid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	_, err := service.DeleteDeviceModel(p)
	if err != nil {
		zap.L().Error(" service.DeleteDeviceModel(p) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}

// CreateDeviceHandler : 删除Device
// @Summary post Device 详情接口
// @Tags Device  相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /device/device/ [post]
func CreateDeviceHandler(c *gin.Context) {
	p := new(models.DeviceBeanParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("CreateDeviceHandler with valid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	_, err := service.CreateDevice(p)
	if err != nil {
		zap.L().Error("service.CreateDevice(p) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}

// GetDeviceHandler : 获取 Device
// @Summary GET Device 详情接口
// @Tags Device  相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /device/device/default/:name/ [get]
func GetDeviceHandler(c *gin.Context) {
	name := c.Param("name")
	device, err := service.GetDevice(name)
	if err != nil {
		zap.L().Error("service.GetDevice(name)", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, device)
}

// DeleteDeviceHandler : 删除 Device
// @Summary Delete Device 详情接口
// @Tags Device  相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /device/device/ [post]
func DeleteDeviceHandler(c *gin.Context) {
	p := new(models.DeviceBeanParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("DeleteDeviceHandler with valid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	_, err := service.DeleteDevice(p)
	if err != nil {
		zap.L().Error("service.DeleteDevice(p) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}

// UpdateDeviceHandler : 更新 Device
// @Summary Post Device 详情接口
// @Tags Device  相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /device/device/ [post]
func UpdateDeviceHandler(c *gin.Context) {
	p := new(models.DeviceUpdateBeanParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("UpdateDeviceHandler with valid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	 err := service.UpdateDevice(p)
	if err != nil {
		zap.L().Error("service.UpdateDevice(p) is failed",zap.Error(err))
		ResponseError(c,CodeServeBusy)
		return
	}
	ResponseSuccess(c,CodeSuccess)
}

