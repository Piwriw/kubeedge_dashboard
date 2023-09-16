package response

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"new-ec-dashboard/models/base"
)

type ConfigMap struct {
	Name       string    `json:"name"`
	Namespace  string    `json:"namespace"`
	DataNum    int       `json:"data_num"`
	CreateTime v1.Time `json:"create_time"`
	// 附件的详情msg
	Data   []base.ListMapItem `json:"data"`
	Labels []base.ListMapItem `json:"labels"`
}
