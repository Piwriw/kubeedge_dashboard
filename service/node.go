package service

import (
	"context"
	"encoding/base64"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"new-ec-dashboard/models"
)

func GetNodeByName(nodeName string) (node *v1.Node, err error) {
	node, err = clientSet.CoreV1().Nodes().Get(context.TODO(), nodeName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return node, nil
}

func GetNodeList(nodeName string) (nodeTemp []v1.Node, err error) {
	nodeList, err := clientSet.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, node := range nodeList.Items {
		if nodeName == "" {
			return nodeList.Items, err
		}
		if nodeName == node.Name {
			nodeTemp := []v1.Node{node}
			return nodeTemp, nil
		}
	}
	return
}

func GetNodeMetrics(nodeName string) (v1.NodeStatus, error) {
	node, err := clientSet.CoreV1().Nodes().Get(context.TODO(), nodeName, metav1.GetOptions{})
	if err != nil {
		return v1.NodeStatus{}, err
	}
	return node.Status, nil
}

func GetJoinToken() (decoder string, err error) {
	secret, _ := clientSet.CoreV1().Secrets("kubeedge").Get(context.TODO(), "tokensecret", metav1.GetOptions{})
	if secret != nil {
		decoder = base64.StdEncoding.EncodeToString(secret.Data["tokendata"])
	} else {
		create, err := clientSet.CoreV1().Secrets("kubeedge").Create(context.TODO(), &v1.Secret{
			TypeMeta: metav1.TypeMeta{
				Kind: "Secret",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "tokensecret",
				Namespace: "kubeedge",
			},
		}, metav1.CreateOptions{})
		if err != nil {
			return "", err
		}
		decoder = base64.StdEncoding.EncodeToString(create.Data["tokendata"])
	}
	return
}

func CreateNodeLabel(nodeL *models.ParamNode) (updateStatus *v1.Node,err error) {
	node, err := clientSet.CoreV1().Nodes().Get(context.TODO(), nodeL.NodeName, metav1.GetOptions{})
	if err != nil {
		return
	}
	labels := node.ObjectMeta.Labels
	labels[nodeL.LabelKey] = labels[nodeL.LabelKey]
	node.Labels = labels
	updateStatus, err = clientSet.CoreV1().Nodes().UpdateStatus(context.TODO(), node, metav1.UpdateOptions{})
	if err != nil {
		return
	}
	return
}

func DeleteNodeLabel(nodeL *models.ParamNode) (update *v1.Node,err error) {
	node, err := clientSet.CoreV1().Nodes().Get(context.TODO(), nodeL.NodeName, metav1.GetOptions{})
	if err != nil {
		return
	}
	labels := node.ObjectMeta.Labels
	delete(labels, nodeL.LabelKey)
	node.ObjectMeta.Labels = labels
	update, err = clientSet.CoreV1().Nodes().Update(context.TODO(), node, metav1.UpdateOptions{})
	if err != nil {
		return
	}
	return
}