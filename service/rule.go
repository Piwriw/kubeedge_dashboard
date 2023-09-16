package service

import (
	"github.com/kubernetes-client/go/kubernetes/client"
	"new-ec-dashboard/models"
)

var (
	routerPointApiVersion = "rules.kubeedge.io/v1"
	routerPointKind       = "RuleEndpoint"
	routerPointGroup      = "rules.kubeedge.io"
	routerPointVersion    = "v1"
	routerPointNamespace  = "default"
	routerPointPlural     = "ruleendpoints"

	ruleGroup      = "rules.kubeedge.io"
	ruleVersion    = "v1"
	ruleNameSpace  = "default"
	rulePlural     = "rules"
	ruleApiVersion = "rules.kubeedge.io/v1"
	ruleKind       = "Rule"
)

func CreateRuleEndPoint(routerParams *models.RouterPointParams) (data interface{}, err error) {
	var routerPointBody models.RouterPointBody
	routerPointBody.ApiVersion = routerPointApiVersion
	routerPointBody.Kind = routerPointKind
	routerPointBody.Metadata.Name = routerParams.Name
	routerPointBody.Spec.RuleEndpointType = routerParams.RuleEndpointType

	if routerParams.Properties == nil || routerParams.Properties == "" {
		routerParams.Properties = "{}"
	}
	routerPointBody.Spec.Properties=routerParams.Properties

	localVarOptionals := new(map[string]interface{})
	data, _, err = clientAPI.CustomObjectsApi.CreateNamespacedCustomObject(ctx, routerPointGroup, routerPointVersion, routerPointNamespace, routerPointPlural, routerPointBody, *localVarOptionals)
	if err != nil {
		return
	}
	return
}

func GetRuleEndPointList() (data interface{}, err error) {
	data, _, err = clientAPI.CustomObjectsApi.GetNamespacedCustomObject(ctx, routerPointGroup, routerPointVersion, routerPointNamespace, routerPointPlural, "")
	if err != nil {
		return
	}
	return
}

func DeleteEndPoint(routerModel *models.RouterPointParams) (data interface{}, err error) {
	m := new(map[string]interface{})
	data, _, err = clientAPI.CustomObjectsApi.DeleteNamespacedCustomObject(ctx, routerPointGroup, routerPointVersion, routerPointNamespace, routerPointPlural, routerModel.Name, client.V1DeleteOptions{}, *m)
	if err != nil {
		return
	}
	return
}

func GetRulePoint(rulePointName string) (data interface{}, err error) {
	data, _, err = clientAPI.CustomObjectsApi.GetNamespacedCustomObject(ctx, routerPointGroup, routerPointVersion, routerPointNamespace, routerPointPlural, rulePointName)
	if err != nil {
		return
	}
	return
}

func CreateRule(ruleModel *models.RuleParams) (data interface{}, err error) {
	var ruleBody models.RuleBody

	ruleBody.ApiVersion = ruleApiVersion
	ruleBody.Kind = ruleKind
	ruleBody.Metadata.Name = ruleModel.Name
	ruleBody.Spec.Source = ruleModel.Source
	ruleBody.Spec.SourceResource = ruleModel.SourceResource
	ruleBody.Spec.Target = ruleModel.Target
	ruleBody.Spec.TargetResource = ruleModel.TargetResource
	localVarOptionals := new(map[string]interface{})
	data, _, err = clientAPI.CustomObjectsApi.CreateNamespacedCustomObject(ctx, ruleGroup, ruleVersion, ruleNameSpace, rulePlural, ruleBody, *localVarOptionals)
	if err != nil {
		return
	}
	return
}

func GetRuleList() (data interface{}, err error) {
	data, _, err = clientAPI.CustomObjectsApi.GetNamespacedCustomObject(ctx, ruleGroup, ruleVersion, ruleNameSpace, rulePlural, "")
	if err != nil {
		return
	}
	return
}

func DeleteRule(ruleModel *models.RuleParams) (data interface{}, err error) {
	m := new(map[string]interface{})
	data, _, err = clientAPI.CustomObjectsApi.DeleteNamespacedCustomObject(ctx, ruleGroup, ruleVersion, ruleNameSpace, rulePlural, ruleModel.Name, client.V1DeleteOptions{}, *m)
	if err != nil {
		return
	}
	return
}

func GetRule(ruleName string) (data interface{}, err error) {
	data, _, err = clientAPI.CustomObjectsApi.GetNamespacedCustomObject(ctx, ruleGroup, ruleVersion, ruleNameSpace, rulePlural, ruleName)
	if err != nil {
		return
	}
	return
}
