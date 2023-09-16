package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"new-ec-dashboard/service"
)

// CreateAIHandler : 创建 AIModel
// @Summary Post AIModel 接口
// @Tags AIModel  相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /ai_model/ai_model/ [post]
//func CreateAIHandler(c *gin.Context) {
//	form, _ := c.MultipartForm()
//	file := form.File["file"][0]
//	content := multipart.GetContent(file)
//	version := form.Value["version"][0]
//	classNames := form.Value["class_names"][0]
//	if classNames == "" || version == "" || file == nil {
//		zap.L().Error("CreateAIHandler 参数缺失")
//		return
//	}
//	err := service.CreateAI(content, version, classNames)
//	if err != nil {
//		zap.L().Error("service.CreateAI(content, version, classNames) is failed", zap.Error(err))
//		ResponseError(c, CodeServeBusy)
//		return
//	}
//	ResponseSuccess(c, CodeSuccess)
//}

// GetListAIHandler : 获得 AIModel 列表
// @Summary Get AIModel 接口
// @Tags AIModel  相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /ai_model/ai_model/ [get]
func GetListAIHandler(c *gin.Context) {
	listAI, err := service.GetListAI()
	if err != nil {
		zap.L().Error("service.GetListAI() is failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, listAI)
}

func DistributeAiHandler(c *gin.Context) {
	// todo 模型下发
}
