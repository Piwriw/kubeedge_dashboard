package service

import (
	"encoding/json"
	"errors"
	"new-ec-dashboard/models"
	"new-ec-dashboard/models/base"
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
	var routerPointBody models.RouterPoint
	routerPointBody.ApiVersion = routerPointApiVersion
	routerPointBody.Kind = routerPointKind
	routerPointBody.Metadata.Name = routerParams.Name
	routerPointBody.Spec.RuleEndpointType = routerParams.RuleEndpointType

	if routerParams.Properties == nil || routerParams.Properties == "" {
		routerParams.Properties = "{}"
	}
	routerPointBody.Spec.Properties = routerParams.Properties

	cr, err := cliCRD.CreateCR(routerPointBody, routerPointGroup, routerPointVersion, routerPointNamespace, routerPointPlural)
	if err != nil {
		return nil, err
	}
	ruleEndPoint := &models.RouterPoint{}
	if err = json.Unmarshal(cr, ruleEndPoint); err != nil {
		return "", err
	}
	return ruleEndPoint, nil
}

func GetRuleEndPointList() (data interface{}, err error) {
	ruleEndpointBytes, err := cliCRD.GetCRList(routerPointGroup, routerPointVersion, routerPointNamespace, routerPointPlural)
	if err != nil {
		return nil, err
	}
	ruleEndPointList := &base.CRDList{}
	if err = json.Unmarshal(ruleEndpointBytes, ruleEndPointList); err != nil {
		return "", err
	}
	return ruleEndPointList, nil
}

func isExistEndPoint(rulePointName string) bool {
	rd, err := GetRulePoint(rulePointName)
	if err != nil {
		return false
	}
	if rd == nil {
		return false
	}
	return true
}
func DeleteEndPoint(routerModel *models.RouterPointParams) error {
	if isExist := isExistEndPoint(routerModel.Name); !isExist {
		return errors.New("不存在这个RuleEndPoint")
	}
	_, err := cliCRD.DeleteCR(routerModel.Name, routerPointGroup, routerPointVersion, routerPointNamespace, routerPointPlural)
	if err != nil {
		return err
	}
	return nil
}

func GetRulePoint(rulePointName string) (data interface{}, err error) {
	cr, err := cliCRD.GetCR(rulePointName, routerPointGroup, routerPointVersion, routerPointNamespace, routerPointPlural)
	if err != nil {
		return nil, err
	}
	ruleEndPoint := &models.RouterPoint{}
	if err = json.Unmarshal(cr, ruleEndPoint); err != nil {
		return "", err
	}
	return ruleEndPoint, nil
}

func CreateRule(ruleModel *models.RuleParams) (data interface{}, err error) {
	var ruleBody models.Rule

	ruleBody.ApiVersion = ruleApiVersion
	ruleBody.Kind = ruleKind
	ruleBody.Metadata.Name = ruleModel.Name
	ruleBody.Spec.Source = ruleModel.Source
	ruleBody.Spec.SourceResource = ruleModel.SourceResource
	ruleBody.Spec.Target = ruleModel.Target
	ruleBody.Spec.TargetResource = ruleModel.TargetResource
	cr, err := cliCRD.CreateCR(ruleBody, ruleGroup, ruleVersion, ruleNameSpace, rulePlural)
	if err != nil {
		return nil, err
	}
	rule := &models.Rule{}
	if err = json.Unmarshal(cr, rule); err != nil {
		return nil, err
	}
	return rule, nil
}

func GetRuleList() (data interface{}, err error) {
	ruleBytes, err := cliCRD.GetCRList(ruleGroup, ruleVersion, ruleNameSpace, rulePlural)
	if err != nil {
		return nil, err
	}
	ruleList := &base.CRDList{}
	if err = json.Unmarshal(ruleBytes, ruleList); err != nil {
		return "", err
	}
	return ruleList, nil
}

func isExistRule(ruleName string) bool {
	rule, err := GetRule(ruleName)
	if err != nil {
		return false
	}
	if rule == nil {
		return false
	}
	return true
}
func DeleteRule(ruleModel *models.RuleParams) error {
	if isExist := isExistRule(ruleModel.Name); !isExist {
		return errors.New("不存在这个Rule")
	}
	_, err := cliCRD.DeleteCR(ruleModel.Name, ruleGroup, ruleVersion, ruleNameSpace, rulePlural)
	if err != nil {
		return err
	}
	return nil
}

func GetRule(ruleName string) (data interface{}, err error) {
	cr, err := cliCRD.GetCR(ruleName, ruleGroup, ruleVersion, ruleNameSpace, rulePlural)
	if err != nil {
		return nil, err
	}
	rule := &models.Rule{}
	if err = json.Unmarshal(cr, rule); err != nil {
		return "", err
	}
	return rule, nil
}
