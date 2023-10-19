package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"new-ec-dashboard/models"
	"new-ec-dashboard/service"
)

// CreateRuleEndPointHandler : 创建 Rule Point
// @Summary Create Rule Point详情接口
// @Tags Rule Point  相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /router/ruleendpoint/ [post]
func CreateRuleEndPointHandler(c *gin.Context) {
	p := new(models.RouterPointParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("CreateRuleEndPointHandler with valid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	point, err := service.CreateRuleEndPoint(p)
	if err != nil {
		zap.L().Error("service.CreateRuleEndPoint(p) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, point)
}

// GetRuleEndPointListHandler : 获取 Rule Point列表
// @Summary Get Rule Point 列表详情接口
// @Tags Rule Point 相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /router/ruleendpoint/ [get]
func GetRuleEndPointListHandler(c *gin.Context) {
	pointList, err := service.GetRuleEndPointList()
	if err != nil {
		zap.L().Error("service.GetRuleEndPointList() is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, pointList)
}

// DeleteEndPointHandler : 删除 Rule Point
// @Summary Delete Rule Point详情接口
// @Tags Rule Point 相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /router/ruleendpoint/ [post]
func DeleteEndPointHandler(c *gin.Context) {
	p := new(models.RouterPointParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("DeleteEndPointHandler with valid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	err := service.DeleteEndPoint(p)
	if err != nil {
		zap.L().Error("service.DeleteEndPoint(p) is failed", zap.Error(err))
		ResponseErrorWithMsg(c, 404, err.Error())
		return
	}
	ResponseSuccess(c, CodeSuccess)
}

// GetRulePointHandler : 获取 Rule Point
// @Summary Get Rule Point详情接口
// @Tags Rule Point 相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /router/ruleendpoint/default/:name/ [get]
func GetRulePointHandler(c *gin.Context) {
	name := c.Param("name")
	rulePoint, err := service.GetRulePoint(name)
	if err != nil {
		zap.L().Error("service.GetRulePoint(name) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, rulePoint)
}

// CreateRuleHandler : 创建 Rule
// @Summary Get Rule 接口
// @Tags Rule  相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /router/rule/ [post]
func CreateRuleHandler(c *gin.Context) {
	p := new(models.RuleParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("CreateRuleHandler with valid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	_, err := service.CreateRule(p)
	if err != nil {
		zap.L().Error(" service.CreateRule(p) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}

// GetRuleListHandler : 获取 Rule
// @Summary Get Rule 接口
// @Tags Rule  相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /router/rule/ [get]
func GetRuleListHandler(c *gin.Context) {
	ruleList, err := service.GetRuleList()
	if err != nil {
		zap.L().Error("service.GetRuleList() is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, ruleList)
}

// DeleteRuleHandler : 删除 Rule
// @Summary Delete Rule 接口
// @Tags Rule  相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /router/rule/ [delete]
func DeleteRuleHandler(c *gin.Context) {
	p := new(models.RuleParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("DeleteRuleHandler with valid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	err := service.DeleteRule(p)
	if err != nil {
		zap.L().Error("service.DeleteRule(p) is failed", zap.Error(err))
		ResponseErrorWithMsg(c, 404, err.Error())
		return
	}
	ResponseSuccess(c, CodeSuccess)
}

// GetRuleHandler : 获取 Rule
// @Summary Get Rule 接口
// @Tags Rule  相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /router/rule/default/:name/ [get]
func GetRuleHandler(c *gin.Context) {
	name := c.Param("name")
	rule, err := service.GetRule(name)
	if err != nil {
		zap.L().Error("service.GetRule(name) is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, rule)
}
