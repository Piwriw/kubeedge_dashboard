package configmap

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"new-ec-dashboard/models/base"
	configmap_req "new-ec-dashboard/models/configmap/request"
)

type CMReq2K8s struct {
}

func (*CMReq2K8s) CmReq2K8s(cmReq configmap_req.ConfigMap) *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cmReq.Name,
			Namespace: cmReq.Namespace,
			Labels:    base.ToMap(cmReq.Labels),
		},
		Data: base.ToMap(cmReq.Data),
	}
}
