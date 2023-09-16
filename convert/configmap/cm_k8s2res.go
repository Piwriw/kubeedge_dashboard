package configmap

import (
	corev1 "k8s.io/api/core/v1"
	"new-ec-dashboard/models/base"
	configmap_res "new-ec-dashboard/models/configmap/response"
)

func GetCMDetailItem(cmK8s corev1.ConfigMap) configmap_res.ConfigMap {
	return configmap_res.ConfigMap{
		Name:       cmK8s.Name,
		Namespace:  cmK8s.Namespace,
		DataNum:    len(cmK8s.Data),
		CreateTime: cmK8s.CreationTimestamp,
	}
}
func GetCMDetailRes(cmK8s corev1.ConfigMap) *configmap_res.ConfigMap {
	detail := GetCMDetailItem(cmK8s)
	detail.Labels = base.ToList(cmK8s.Labels)
	detail.Data = base.ToList(cmK8s.Data)
	return &detail
}
