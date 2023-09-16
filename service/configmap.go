package service

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	cm_convert"new-ec-dashboard/convert/configmap"
	configmap_res "new-ec-dashboard/models/configmap/response"
	"strings"
)
func  GetConfigMapDetail(namespace, name string) (*configmap_res.ConfigMap, error) {
	k8sCm, err :=clientSet.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return cm_convert.GetCMDetailRes(*k8sCm), nil
}
func  GetConfigMapList(namespace, keyword string) ([]configmap_res.ConfigMap, error) {
	configMapList, err := clientSet.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	cmListRes := make([]configmap_res.ConfigMap, 0)
	for _, item := range configMapList.Items {
		if !strings.Contains(item.Name, keyword) {
			continue
		}
		cmListRes = append(cmListRes, cm_convert.GetCMDetailItem(item))
	}
	return cmListRes, nil
}