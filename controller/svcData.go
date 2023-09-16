package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"new-ec-dashboard/models"
	"new-ec-dashboard/service"
)
// PostSvcDataListHandler : 存储 SvcData
// @Summary Post SvcData 接口
// @Tags SvcData  相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /svc_data/svc_data/ [Post]
func PostSvcDataListHandler(c *gin.Context) {
	p := new(models.SvcDataParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("PostSvcDataListHandler with valid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	err := service.InsertSvcData(p)
	if err != nil {
		zap.L().Error("service.InsertSvcData(p) is failed",zap.Error(err))
		ResponseError(c,CodeServeBusy)
		return
	}
	ResponseSuccess(c,CodeSuccess)
}

// GetSvcDataListHandler : 获取 SvcData
// @Summary Get SvcData 接口
// @Tags SvcData  相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Router /svc_data/svc_data/ [get]
func GetSvcDataListHandler(c *gin.Context) {
	svcDataList, err := service.GetSvcDataList()

	if err != nil {
		zap.L().Error("service.GetSvcDataList() is failed",zap.Error(err))
		ResponseError(c,CodeServeBusy)
		return
	}
	ResponseSuccess(c,svcDataList)
}

func PostIeDataHandler(c *gin.Context) {
	p := new(models.IeParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("PostIeDataHandler with valid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	err := service.InsertIeData(p)
	if err != nil {
		zap.L().Error("service.InsertIeData(p) is failed",zap.Error(err))
		ResponseError(c,CodeServeBusy)
		return
	}
	ResponseSuccess(c,CodeSuccess)
}
func GetIeDataListHandler(c *gin.Context) {
	pagenum, pagesize := getPageInfo(c)
	page, err := service.GetIeDataList(pagenum,pagesize)
	if err != nil {
		zap.L().Error("service.GetSvcDataList() is failed",zap.Error(err))
		ResponseError(c,CodeServeBusy)
		return
	}
	ResponseSuccess(c,page)
}

func GetSvcDataHandler(c *gin.Context){

}