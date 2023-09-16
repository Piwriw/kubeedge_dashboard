package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"io/ioutil"
	"new-ec-dashboard/service"
)

// GetPodListHandler : 获取Pod列表
// @Summary Get Pod接口
// @Tags Pod相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /pod/pod/default/ [get]
func GetPodListHandler(c *gin.Context) {
	podList, err := service.PodFetchList()
	if err != nil {
		zap.L().Error("service.PodFetchList() is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, podList)
}

// GetPodHandler : 获取Pod详情
// @Summary Get Pod接口
// @Tags Pod相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /pod/pod/:namespace/:name/ [get]
func GetPodHandler(c *gin.Context) {
	name := c.Param("name")
	namespace := c.Param("namespace")
	pod, err := service.GetPodByNameDefault(name, namespace)
	if err != nil {
		zap.L().Error("service.GetPodByNameDefault(name, namespace) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, pod)
}


// CreateAPPHandler : 创建Pod详情
// @Summary Create Pod接口
// @Tags Pod相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /pod/ [post]
func CreateAPPHandler(c *gin.Context) {
	conf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		zap.L().Error("ioutil.ReadAll(c.Request.Body) is failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	app, err := service.CreateAPP(conf)
	if err != nil {
		zap.L().Error(" service.CreateAPP(conf) is failed", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	ResponseSuccess(c,app)
}


// DeletePodHandler : 删除Pod详情
// @Summary Delete Pod接口
// @Tags Pod相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /:namespace/:name/ [post]
func DeletePodHandler(c *gin.Context) {
	nameSpace:= c.Param("namespace")
	name:= c.Param("name")
	flag, err := service.DeletePod(name, nameSpace)
	if err != nil {
		zap.L().Error("service.DeletePod(name, nameSpace) is failed",zap.Error(err))
		ResponseError(c,CodeServeBusy)
		return
	}
	ResponseSuccess(c,flag)
}
