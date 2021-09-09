package resources

import (
	"context"
	"github.com/bloveing/kube-client/pkg/client"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNodeList() (*apiv1.NodeList, error) {
	clientSet, err := client.ClientSet()
	if err != nil {
		return nil, err
	}

	config, err := clientSet.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	return config, err
}

func GetNodeByName(name string) (*apiv1.Node, error) {
	clientSet, err := client.ClientSet()
	if err != nil {
		return nil, err
	}

	config, err := clientSet.CoreV1().Nodes().Get(context.Background(), name, metav1.GetOptions{})
	return config, err
}

func DeleteNode(name string) error {
	clientSet, err := client.ClientSet()
	if err != nil {
		return err
	}

	err = clientSet.CoreV1().Nodes().Delete(context.Background(), name, metav1.DeleteOptions{})
	return err
}

// TODO: AddNodeLabel
//func AddNodeLabel(data models.NodeLabels) (*apiv1.Node, error) {
//
//}
// TODO: DeleteNodeLabel
//func DeleteNodeLabel(nodename, key string) (*apiv1.Node, error) {
//
//}

// TODO: AddNodeTaint
// TODO: drainNode
