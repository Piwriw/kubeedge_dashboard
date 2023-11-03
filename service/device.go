package service

import (
	"encoding/json"
	"errors"
	"new-ec-dashboard/models"
	"new-ec-dashboard/models/base"
)

const (
	DeviceModelGroup      = "devices.kubeedge.io"
	DeviceModelVersion    = "v1alpha2"
	DeviceModelNamespace  = "default"
	DeviceModelPlural     = "devicemodels"
	DeviceModelAPIVersion = "devices.kubeedge.io/v1alpha2"
	DeviceModelKind       = "DeviceModel"

	DeviceGroup      = "devices.kubeedge.io"
	DeviceVersion    = "v1alpha2"
	DeviceNameSpace  = "default"
	DevicePlural     = "devices"
	DeviceApiVersion = "devices.kubeedge.io/v1alpha2"
	DeviceKind       = "Device"
)

func CreateDevice(deviceBean *models.DeviceBeanParams) error {
	exist := isDeviceExist(deviceBean.Name)
	if exist {
		return errors.New("已经存在同名Device")
	}
	modelExist := isDeviceModelExist(deviceBean.DMName)
	if !modelExist {
		return errors.New("不存在这个DeviceModel")
	}
	deviceModel, err := GetDeviceModel(deviceBean.DMName)
	if err != nil {
		return err
	}

	device := &models.Device{}
	device.ApiVersion = DeviceApiVersion
	device.Kind = DeviceKind
	device.Metadata.Name = deviceBean.Name
	device.Spec.DeviceModelRef.Name = deviceBean.DMName

	matchExpressions := []models.MatchExpressionsBody{{Key: "", Operator: "In", Values: []string{deviceBean.NodeName}}}
	node := []models.MatchExpressions{{matchExpressions}}
	device.Spec.NodeSelector.NodeSelectorTerms = node

	var twins []models.DeviceTwins
	for _, property := range deviceModel.Spec.Properties {
		var twin models.DeviceTwins
		twin.PropertyName = property.Name
		twin.Desired.Metadata.Type = "string"
		twin.Desired.Value = ""
		twin.Reported.Metadata.Type = "string"
		twin.Reported.Value = ""
		twins = append(twins, twin)
	}
	device.Status.Twins = twins
	_, err = cliCRD.CreateCR(device, DeviceGroup, DeviceVersion, DeviceNameSpace, DevicePlural)
	if err != nil {
		return err
	}
	return nil
}

func GetDevice(deviceName string) (*models.Device, error) {
	cr, err := cliCRD.GetCR(deviceName, DeviceGroup, DeviceVersion, DeviceNameSpace, DevicePlural)
	if err != nil {
		return nil, err
	}
	device := &models.Device{}
	if err = json.Unmarshal(cr, device); err != nil {
		return nil, err
	}
	return device, nil
}

func isDeviceExist(deviceName string) bool {
	device, err := GetDevice(deviceName)
	if err != nil {
		return false
	}
	if device.Spec.DeviceModelRef.Name == "" {
		return false
	}
	return true
}

func DeleteDevice(deviceBean *models.DeviceBeanParams) (err error) {
	if isExist := isDeviceExist(deviceBean.Name); !isExist {
		return errors.New("不存在这个Device")
	}
	_, err = cliCRD.DeleteCR(deviceBean.Name, DeviceGroup, DeviceVersion, DeviceNameSpace, DevicePlural)
	if err != nil {
		return err
	}
	return nil
}

func UpdateDevice(modelBean *models.DeviceUpdateBeanParams) (b bool, err error) {
	if exist := isDeviceExist(modelBean.Name); !exist {
		return false, errors.New("不存在这个Device")
	}
	deviceOld, err := GetDevice(modelBean.Name)
	if err != nil {
		return false, err
	}
	newTwins := make([]models.DeviceTwins, 0)
	for _, twin := range deviceOld.Status.Twins {
		m := modelBean.Desired.(map[string]interface{})
		for key, val := range m {
			if key == twin.PropertyName {
				twin.Desired.Value = val.(string)
			}
		}

		newTwins = append(newTwins, twin)
	}
	deviceOld.Status.Twins = newTwins
	// 更新状态
	ok, err := cliCRD.UpdateCR(deviceOld, DeviceGroup, DeviceVersion, DeviceNameSpace, DevicePlural, deviceOld.Metadata.Name)
	if err != nil {
		return false, err
	}
	if ok == nil {
		return false, errors.New("更新失败")
	}
	return true, nil
}

func GetDeviceList() (data interface{}, err error) {
	crds, err := cliCRD.GetCRList(DeviceGroup, DeviceVersion, DeviceNameSpace, DevicePlural)
	if err != nil {
		return nil, err
	}
	deviceList := &base.CRDList{}
	if err = json.Unmarshal(crds, deviceList); err != nil {
		return "", err
	}
	return deviceList, nil
}

func CreateDeviceModel(deviceModelParam *models.DeviceModelParam) error {
	var properties []models.PropertiesBody
	if exist := isDeviceModelExist(deviceModelParam.Name); exist {
		return errors.New("已经存在这个DeviceModel")
	}
	for _, val := range deviceModelParam.Properties {
		var p models.PropertiesBody
		p.Name = val
		p.Type.String.AccessMode = "ReadWrite"
		p.Type.String.DefaultValue = ""
		properties = append(properties, p)
	}
	var body models.DeviceModel
	body.ApiVersion = DeviceModelAPIVersion
	body.Kind = DeviceModelKind
	body.Metadata.Name = deviceModelParam.Name
	body.Spec.Properties = properties

	_, err := cliCRD.CreateCR(body, DeviceModelGroup, DeviceModelVersion, DeviceModelNamespace, DeviceModelPlural)
	if err != nil {
		return nil
	}
	return nil
}

func GetDeviceModelList() (data interface{}, err error) {
	crs, err := cliCRD.GetCRList(DeviceModelGroup, DeviceModelVersion, DeviceModelNamespace, DeviceModelPlural)
	if err != nil {
		return nil, err
	}
	deviceModelList := &base.CRDList{}
	if err = json.Unmarshal(crs, deviceModelList); err != nil {
		return nil, err
	}
	return deviceModelList, nil
}

func GetDeviceModel(modelName string) (data *models.DeviceModel, err error) {
	cr, err := cliCRD.GetCR(modelName, DeviceModelGroup, DeviceModelVersion, DeviceModelNamespace, DeviceModelPlural)
	deviceModel := &models.DeviceModel{}
	if err = json.Unmarshal(cr, deviceModel); err != nil {
		return nil, err
	}
	return deviceModel, nil
}

func UpdateDeviceModel(deviceModelParam *models.DeviceModelParam) (data interface{}, err error) {
	deviceModel, err := GetDeviceModel(deviceModelParam.Name)
	if err != nil {
		return
	}
	propertiesBody := make([]models.PropertiesBody, 0)
	for _, property := range deviceModelParam.Properties {
		p := &models.PropertiesBody{Name: property}
		p.Type.String.AccessMode = "ReadWrite"
		p.Type.String.DefaultValue = ""
		propertiesBody = append(propertiesBody, *p)
	}
	deviceModel.Spec.Properties = propertiesBody
	_, err = cliCRD.UpdateCR(deviceModel, DeviceModelGroup, DeviceModelVersion, DeviceModelNamespace, DeviceModelPlural, deviceModelParam.Name)
	if err != nil {
		return false, err
	}
	return true, nil
}

func isDeviceModelExist(deviceModelName string) bool {
	devicemodel, err := GetDeviceModel(deviceModelName)
	if err != nil {
		return false
	}
	if devicemodel.Spec.Properties == nil {
		return false
	}
	return true
}

func DeleteDeviceModel(deviceModel *models.DeviceModelParam) error {
	if exist := isDeviceModelExist(deviceModel.Name); !exist {
		return errors.New("不存在这个DeviceModel")
	}
	_, err := cliCRD.DeleteCR(deviceModel.Name, DeviceModelGroup, DeviceModelVersion, DeviceModelNamespace, DeviceModelPlural)
	if err != nil {
		return err
	}
	return nil
}
