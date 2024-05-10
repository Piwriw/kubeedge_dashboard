package service

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"new-ec-dashboard/setting"
	"new-ec-dashboard/utils"
)

var clientSet *kubernetes.Clientset
var ctx context.Context
var cliCRD *utils.CRDClient

func InitK8sClient() error {
	var err error
	switch setting.Conf.K8sConfig.AuthModel {
	case "token":
		err = ClientWithToken()
	case "config":
		err = ClientWithConfig()
	default:
		return errors.New("目前只支持 token | config")
	}
	if err != nil {
		return err
	}
	return nil
}
func ClientWithConfig() error {
	kubeconifg := "config"
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconifg)
	if err != nil {
		return err
	}

	// create the clientset
	clientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	ctx = context.Background()
	return nil
}

func ClientWithToken() error {
	// init CRDClient
	cliCRD = utils.CreateCRDClient(setting.Conf.K8sConfig.Token, setting.Conf.K8sConfig.Host, "6443")

	config := &rest.Config{
		Host:        fmt.Sprintf("https://%s:6443", setting.Conf.K8sConfig.Host),
		BearerToken: setting.Conf.K8sConfig.Token,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true, // 设置为true时 不需要CA
			CAData:   []byte(""),
		},
	}
	clientSet2, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	clientSet = clientSet2
	return nil
}

func GetClientSet() *kubernetes.Clientset {
	return clientSet
}
