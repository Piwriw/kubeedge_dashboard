package models

import "time"

// 定义请求参数结构体

// ParamLogin :登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamNode : Node请求参数
type ParamNode struct {
	NodeName   string `json:"node_name" binding:"required"`
	LabelKey   string `json:"label_key" `
	LabelValue string `json:"label_value" `
}

// DeviceModelParam : 设备模板请求参数
type DeviceModelParam struct {
	Name       string   `json:"name" binding:"required"`
	Properties []string `json:"properties"`
}

// DeviceBeanParams : 设备实例去请求参数
type DeviceBeanParams struct {
	DmName   string `json:"dm_name" `
	Name     string `json:"name" `
	NodeName string `json:"node_name" `
}

type DeviceUpdateBeanParams struct {
	Name    string      `json:"name" binding:"required"`
	Desired interface{} `json:"desired" binding:"required"`
}

// RouterPointParams : 路由节点请求参数
type RouterPointParams struct {
	Name             string      `json:"name" binding:"required"`
	RuleEndpointType string      `json:"rule_endpoint_type"`
	Properties       interface{} `json:"properties" `
}

// RuleParams : 路由请求参数
type RuleParams struct {
	Name           string      `json:"name" binding:"required"`
	Source         string      `json:"source" `
	SourceResource interface{} `json:"sourceResource"`
	Target         string      `json:"target"`
	TargetResource interface{} `json:"targetResource"`
}

// SvcDataParams : 数据请求参数
type SvcDataParams struct {
	ImageBase64 string    `json:"image_base64" db:"image_base64"  binding:"required"`
	ClassName   string    `json:"class_name" db:"class_name"  binding:"required"`
	UpLoadTime  time.Time `json:"upload_time" db:"upload_time" `
}

// IeParams : 电流请求参数
type IeParams struct {
	EdgeDeviceId           string  `json:"edge_device_id" db:"edge_device_id"`
	Voltage                float32 `json:"voltage" db:"voltage"`
	ElectricCurrent        float32 `json:"electric_current" db:"electric_current"`
	Power                  float32 `json:"power" db:"power"`
	PowerFactor            float32 `json:"power_factor" db:"power_factor"`
	Frequency              float32 `json:"frequency" db:"frequency"`
	Label                  string  `json:"label" db:"label"`
	ElectricityConsumption float32 `json:"electricity_consumption" db:"electricity_consumption"`
}
