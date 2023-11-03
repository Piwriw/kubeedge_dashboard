package service

import (
	"context"
	"errors"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
)

var (
	ErrorEmptyYaml   = errors.New("Conf YAML is Empty !")
	ErrorInvalidYaml = errors.New("Conf YAML with Invalid Params!")
)

func PodFetchList() (podList *v1.PodList, err error) {
	podList, err = clientSet.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return
	}
	return
}

func GetPodByNameDefault(name string, namespace string) (pod *v1.Pod, err error) {
	pod, err = clientSet.CoreV1().Pods(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return
	}
	return
}

func CreateAPP(conf []byte) (data v1.Pod, err error) {
	if len(conf) == 0 {
		err = ErrorEmptyYaml
		return
	}
	podBody := new(v1.Pod)
	err = yaml.Unmarshal(conf, &podBody)
	if err != nil {
		err = ErrorInvalidYaml
		return
	}

	_, err = clientSet.CoreV1().Pods(podBody.Namespace).Create(context.TODO(), podBody, metav1.CreateOptions{})
	if err != nil {
		return
	}
	return
}

func DeletePod(name, namespace string) (flag bool, err error) {
	err = clientSet.CoreV1().Pods(namespace).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		flag = false
		return
	}
	flag = true
	return
}
