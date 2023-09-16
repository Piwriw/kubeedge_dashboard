package service

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	secret_convert"new-ec-dashboard/convert/secret"
	secret_res "new-ec-dashboard/models/secret/response"
	"strings"
)
func  GetSecretDetail(namespace, name string) (*secret_res.Secret, error) {
	secretK8s, err :=clientSet.CoreV1().Secrets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	secretRes := secret_convert.SecretK8s2ResDetail(*secretK8s)
	return &secretRes, err
}
func  GetSecretList(namespace, keyword string) ([]secret_res.Secret, error) {
	list, err := clientSet.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	secretResList := make([]secret_res.Secret, 0)
	for _, item := range list.Items {
		if !strings.Contains(item.Name, keyword) {
			continue
		}
		secretResList = append(secretResList, secret_convert.SecretK8s2ResItem(item))
	}
	return secretResList, nil
}