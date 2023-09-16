package response

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"new-ec-dashboard/models/base"
)

type Secret struct {
	Name       string             `json:"name"`
	Namespace  string             `json:"namespace"`
	DataNum    int                `json:"data_num"`
	CreateTime v1.Time            `json:"create_time"`
	Type       corev1.SecretType  `json:"type"`
	Labels     []base.ListMapItem `json:"labels"`
	Data       []base.ListMapItem `json:"data"`
}
