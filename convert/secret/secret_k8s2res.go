package secret

import (
	corev1 "k8s.io/api/core/v1"
	"new-ec-dashboard/models/base"
	secret_res "new-ec-dashboard/models/secret/response"
)

func  SecretK8s2ResItem(secretK8s corev1.Secret) secret_res.Secret {
	return secret_res.Secret{
		Name:      secretK8s.Name,
		Namespace: secretK8s.Namespace,
		Type:      secretK8s.Type,
		DataNum:   len(secretK8s.StringData),
		CreateTime:       secretK8s.CreationTimestamp,
	}
}
func  SecretK8s2ResDetail(secretK8s corev1.Secret) secret_res.Secret {
	return secret_res.Secret{
		Name:      secretK8s.Name,
		Namespace: secretK8s.Namespace,
		Type:      secretK8s.Type,
		Data:      base.ToListWithMapByte(secretK8s.Data),
		Labels:    base.ToList(secretK8s.Labels),
	}
}
