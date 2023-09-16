package router

import (
	"github.com/gin-gonic/gin"
	"new-ec-dashboard/controller"
)

func RuleRouter(r *gin.Engine) {
	ruleGroup := r.Group("/router")
	{
		ruleGroup.POST("/ruleendpoint/", controller.CreateRuleEndPointHandler)
		ruleGroup.GET("/ruleendpoint/", controller.GetRuleEndPointListHandler)
		ruleGroup.DELETE("/ruleendpoint/", controller.DeleteEndPointHandler)
		ruleGroup.GET("/ruleendpoint/default/:name/", controller.GetRulePointHandler)
		ruleGroup.POST("/rule/", controller.CreateRuleHandler)
		ruleGroup.GET("/rule/", controller.GetRuleListHandler)
		ruleGroup.DELETE("/rule/", controller.DeleteRuleHandler)
		ruleGroup.GET("/rule/default/:name/", controller.GetRuleHandler)
	}

}