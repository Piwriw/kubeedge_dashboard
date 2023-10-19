package service

import (
	"context"
	"crypto/tls"
	"errors"
	"flag"
	"fmt"
	"github.com/kubernetes-client/go/kubernetes/client"
	cf "github.com/kubernetes-client/go/kubernetes/config"
	"io/ioutil"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"net/http"
	"new-ec-dashboard/setting"
	"new-ec-dashboard/utils"
	"path/filepath"
)

var clientSet *kubernetes.Clientset
var clientAPI *client.APIClient
var ctx context.Context
var cliCRD *utils.CRDClient

//type ApiKeyBody struct {
//	APIKey client.APIKey
//}

func InitK8SLinux() {

	//Linux
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join("../", ".kube", "config"))
	if err != nil {
		panic(err)
	}
	//创建clientset
	clientSet, err = kubernetes.NewForConfig(config)
	//restClient, err := rest.RESTClientFor(config)

	if err != nil {
		panic(err)
	}

	//token, err := ioutil.ReadFile("D:\\Project\\ec-dashboard\\config\\k8sconfig\\token.txt")
	token, err := ioutil.ReadFile("./conf/k8sconfig/token.txt")
	filepath.Join()
	if err != nil {
		panic(err)
	}
	//ca, err := ioutil.ReadFile("D:\\Project\\ec-dashboard\\config\\k8sconfig\\ca.crt")
	if err != nil {
		panic(err)
	}
	configLoader, err := cf.NewKubeConfigLoaderFromYAMLFile(filepath.Join("../", ".kube", "config"), false)
	//configLoader, err := cf.NewKubeConfigLoaderFromYAMLFile("F:\\project\\Intelligent_manangement_platform\\new-ec-dashboard\\conf\\k8sconfig\\kubeconfig", false)

	configuration, err := configLoader.LoadAndSet()

	ctx = context.Background()
	ctx = context.WithValue(ctx, client.ContextAPIKey, client.APIKey{Key: "Bearer " + string(token), Prefix: ""})

	clientAPI = client.NewAPIClient(configuration)
	//configuration := client.NewConfiguration()

	//configLoader, err := cf.NewKubeConfigLoaderFromYAMLFile("D:\\Project\\ec-dashboard\\config\\k8sconfig\\kubeconfig", false)
	//configLoader, err := cf.NewKubeConfigLoaderFromYAMLFile("./kubeconfig", false)
	//
	//configuration, err := configLoader.LoadAndSet()
	//
	//ctx = context.Background()
	//ctx = context.WithValue(ctx, client.ContextAPIKey, client.APIKey{Key: "Bearer " + string(token), Prefix: ""})
	//
	//
	//clientAPI = client.NewAPIClient(configuration)

}
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

	configLoader, err := cf.NewKubeConfigLoaderFromYAMLFile(kubeconifg, false)

	configuration, err := configLoader.LoadAndSet()

	ctx = context.Background()
	//ctx = context.WithValue(ctx, client.ContextAPIKey, client.APIKey{Key: "Bearer " + string(token), Prefix: ""})
	clientAPI = client.NewAPIClient(configuration)
	return nil
}

func ClientWithToken() error {
	// init CRDClient
	cliCRD = utils.CreateCRDClient(setting.Conf.K8sConfig.Token, setting.Conf.K8sConfig.Host, "6443")

	token, err := ioutil.ReadFile(setting.Conf.K8sConfig.Token)
	if err != nil {
		return err
	}

	config := &rest.Config{
		Host:        fmt.Sprintf("https://%s:6443", setting.Conf.K8sConfig.Host),
		BearerToken: string(token),
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true, // 设置为true时 不需要CA
			CAData:   []byte(""),
		},
	}
	clientSet2, err := kubernetes.NewForConfig(config)
	clientSet = clientSet2
	if err != nil {
		return err
	}
	configuration := &client.Configuration{
		BasePath: fmt.Sprintf("https://%s:6443", setting.Conf.K8sConfig.Host),
		DefaultHeader: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", string(token)),
		},
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
	}
	clientAPI = client.NewAPIClient(configuration)
	return nil
}

func GetClientSet() *kubernetes.Clientset {
	return clientSet
}

// GetClientAPI : 获取自定义资源 -k8s client
func GetClientAPI() *client.APIClient {
	return clientAPI
}
func GetContent() context.Context {
	return ctx
}
