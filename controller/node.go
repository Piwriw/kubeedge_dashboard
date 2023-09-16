package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"new-ec-dashboard/models"
	"new-ec-dashboard/service"
)

// GetNodeByNameHandler : 获取Node信息
// @Summary 获取Node接口
// @Description 获取Node信息
// @Tags Node相关接口
// @Accept application/json
// @Produce application/json
// Param: nodename string
// @Security ApiKeyAuth
// @Router /node/node/:nodename/: [get]
func GetNodeByNameHandler(c *gin.Context) {
	nodeName := c.Param("nodename")

	node, err := service.GetNodeByName(nodeName)
	if err != nil {
		zap.L().Error("service.GetNodeByName failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, node)
}

// GetNodeListHandler : 获取Node列表信息
// @Summary 获取Node接口
// @Description 获取Node信息
// @Tags Node相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /node/node/: [get]
func GetNodeListHandler(c *gin.Context) {
	name := c.Query("name")
	nodeList, err := service.GetNodeList(name)
	if err != nil {
		zap.L().Error("service.GetNodeList(name) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, nodeList)
}

// GetNodeMetricsHandler : 获取Node运行状态信息
// @Summary 获取Node接口
// @Description 获取Node信息
// @Tags Node相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /node/metrics/:nodename/: [get]
func GetNodeMetricsHandler(c *gin.Context) {
	nodename := c.Param("nodename")
	metrics, err := service.GetNodeMetrics(nodename)
	if err != nil {
		zap.L().Error("service.GetNodeMetrics(nodename) failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, metrics)
}

// GetJoinTokenHandler : 获取加入Node的Token
// @Summary 获取Node接口
// @Description 获取Node信息
// @Tags Node相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /node/join/ [get]
func GetJoinTokenHandler(c *gin.Context) {
	joinToken, err := service.GetJoinToken()
	if err != nil {
		zap.L().Error("service.GetJoinToken() failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, joinToken)
}

// CreateNodeLabelHandler : 创建 Node 标签
// @Summary Create Node接口
// @Description 获取Node信息
// @Tags Node相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /node/label/ [put]
func CreateNodeLabelHandler(c *gin.Context) {
	p := new(models.ParamNode)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("CreateNodeLabel with valid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	label, err := service.CreateNodeLabel(p)
	if err != nil {
		zap.L().Error("service.CreateNodeLabel(p)", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, label)
}

// DeleteNodeLabelHandler : 删除 Node 标签
// @Summary Delete Node接口
// @Description Delete Node信息
// @Tags Node相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /node/label/ [post]
func DeleteNodeLabelHandler(c *gin.Context) {
	p := new(models.ParamNode)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("DeleteNodeLabelHandler with valid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	label, err := service.DeleteNodeLabel(p)
	if err != nil {
		zap.L().Error("service.DeleteNodeLabel(p)",zap.Error(err))
		ResponseError(c,CodeServeBusy)
		return
	}
	ResponseSuccess(c,label)
}
