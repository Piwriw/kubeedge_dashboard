package service

import (
	"encoding/json"
	"github.com/kubernetes-client/go/kubernetes/client"
	"new-ec-dashboard/models"
)

var (
	group     = "devices.kubeedge.io"
	version   = "v1alpha2"
	namespace = "default"
	plural    = "devicemodels"

	deviceGroup      = "devices.kubeedge.io"
	deviceVersion    = "v1alpha2"
	deviceNameSpace  = "default"
	devicePlural     = "devices"
	deviceApiVersion = "devices.kubeedge.io/v1alpha2"
	deviceKind       = "Device"
)

func GetDeviceList() (data interface{}, err error) {
	data, _, err = clientAPI.CustomObjectsApi.GetNamespacedCustomObject(ctx, deviceGroup, deviceVersion, deviceNameSpace, devicePlural, "")
	if err != nil {
		return
	}
	return
}

func CreateDeviceModel(deviceModelParam *models.DeviceModelParam) (data interface{},err error) {
	var properties []models.PropertiesBody
	for _, val := range deviceModelParam.Properties {
		var p models.PropertiesBody
		p.Name = val
		p.Type.String.AccessMode = "ReadWrite"
		p.Type.String.DefaultValue = ""
		properties = append(properties, p)
	}
	var body models.DeviceBody
	body.ApiVersion = "devices.kubeedge.io/v1alpha2"
	body.Kind = "DeviceModel"
	body.Metadata.Name = deviceModelParam.Name
	body.Spec.Properties = properties
	localVarOptionals := new(map[string]interface{})
	//bodyJSON, err := json.Marshal(properties)
	data, _, err = clientAPI.CustomObjectsApi.CreateNamespacedCustomObject(ctx, group, version, namespace, plural, body, *localVarOptionals)
	if err != nil {
		return
	}
	return
}

func GetDeviceModelList() (data interface{},err error) {
	data, _, err = clientAPI.CustomObjectsApi.GetNamespacedCustomObject(ctx, group, version, namespace, plural, "")
	if err != nil {
		return
	}
	return
}

func GetDeviceModel(modelName string) (data interface{},err error) {
	data, _, err = clientAPI.CustomObjectsApi.GetNamespacedCustomObject(ctx, group, version, namespace, plural, modelName)
	if err != nil {
		return
	}
	return
}

func UpdateDeviceModel(deviceModel *models.DeviceModelParam) (data interface{},err error) {
	oldModel, _, err := clientAPI.CustomObjectsApi.GetNamespacedCustomObject(ctx, group, version, namespace, plural, deviceModel.Name)
	if err != nil {
		return
	}
	var properties []models.PropertiesBody
	for _, val := range deviceModel.Properties {
		var p models.PropertiesBody
		p.Name = val
		p.Type.String.AccessMode = "ReadWrite"
		p.Type.String.DefaultValue = ""
		properties = append(properties, p)
	}

	m := oldModel.(map[string]interface{})
	spec := m["spec"]
	spec.(map[string]interface{})["properties"] = properties
	data, _, err = clientAPI.CustomObjectsApi.ReplaceNamespacedCustomObject(ctx, group, version, namespace, plural, deviceModel.Name, oldModel)
	if err != nil {
		return
	}
	return
}

func DeleteDeviceModel(deviceModel *models.DeviceModelParam) (data interface{},err error) {
	m := new(map[string]interface{})
	data, _, err = clientAPI.CustomObjectsApi.DeleteNamespacedCustomObject(ctx, group, version, namespace, plural, deviceModel.Name, client.V1DeleteOptions{}, *m)
	if err != nil {
		return
	}
	return
}

func CreateDevice(deviceBean *models.DeviceBeanParams) (data interface{},err error) {
	localVarOptionals := new(map[string]interface{})
	deviceModel, _, err := clientAPI.CustomObjectsApi.GetNamespacedCustomObject(ctx, group, version, namespace, plural, deviceBean.DmName)
	if err != nil {
		return
	}

	m := deviceModel.(map[string]interface{})
	spec := m["spec"]
	p := spec.(map[string]interface{})["properties"]
	properties := p.([]interface{})
	var twins []models.DeviceTwins
	for _, val := range properties {
		tmp := val.(map[string]interface{})
		var twin models.DeviceTwins
		twin.PropertyName = tmp["name"].(string)
		twin.Desired.Metadata.Type = "string"
		twin.Desired.Value = ""
		twin.Reported.Metadata.Type = "string"
		twin.Reported.Value = ""
		twins = append(twins, twin)
	}
	var deviceBeanBody models.DeviceBeanBody
	deviceBeanBody.ApiVersion = deviceApiVersion
	deviceBeanBody.Kind = deviceKind
	deviceBeanBody.Metadata.Name = deviceBean.Name
	deviceBeanBody.Spec.DeviceModelRef.Name = deviceBean.DmName

	matchExpressions := []models.MatchExpressionsBody{{Key: "", Operator: "In", Values: []string{deviceBean.NodeName}}}
	node := []models.MatchExpressions{{matchExpressions}}
	deviceBeanBody.Spec.NodeSelector.NodeSelectorTerms = node

	deviceBeanBody.Status.Twins = twins

	data, _, err = clientAPI.CustomObjectsApi.CreateNamespacedCustomObject(ctx, deviceGroup, deviceVersion, deviceNameSpace, devicePlural, deviceBeanBody, *localVarOptionals)
	if err != nil {
		return
	}
	return
}

func GetDevice(deviceName string) (data interface{},err error) {
	data, _, err = clientAPI.CustomObjectsApi.GetNamespacedCustomObject(ctx, deviceGroup, deviceVersion, deviceNameSpace, devicePlural, deviceName)
	if err != nil {
		return
	}
	return
}

func DeleteDevice(deviceBean *models.DeviceBeanParams) (data interface{},err error) {
	localVarOptionals := new(map[string]interface{})
	data, _, err = clientAPI.CustomObjectsApi.DeleteNamespacedCustomObject(ctx, deviceGroup, deviceVersion, deviceNameSpace, devicePlural, deviceBean.Name, client.V1DeleteOptions{}, *localVarOptionals)
	if err != nil {
		return
	}
	return
}

func UpdateDevice(modelBean *models.DeviceUpdateBeanParams) (err error){
	oldData, _, err := clientAPI.CustomObjectsApi.GetNamespacedCustomObject(ctx, deviceGroup, deviceVersion, deviceNameSpace, devicePlural, modelBean.Name)
	if err != nil {
		return
	}
	// 依次获取到字段属性
	twinsT := oldData.(map[string]interface{})["status"].(map[string]interface{})["twins"].([]interface{})
	var newTwins []models.DeviceTwins
	for _, v :=range twinsT {
		var dt models.DeviceTwins
		bytes, err := json.Marshal(v)
		if err != nil {
			return err
		}
		err = json.Unmarshal(bytes, &dt)
		if err != nil {
			return err
		}
		// 给更新的字段属性更新
		for key,dev:= range modelBean.Desired.(map[string]interface{}){
			propertyName := dt.PropertyName
			if key==propertyName {
				dt.Desired.Value=dev.(string)
			}
		}
		newTwins=append(newTwins,dt)
	}
	oldData.(map[string]interface{})["status"].(map[string]interface{})["twins"]=newTwins
	// 在通过api-server 更新集群中YAML文件
	_, _, err = clientAPI.CustomObjectsApi.ReplaceNamespacedCustomObject(ctx, deviceGroup, deviceVersion, deviceNameSpace, devicePlural, modelBean.Name , oldData)
	if err != nil {
		return
	}
	return
}