package resources

import (
	"context"
	"github.com/bloveing/kube-client/pkg/client"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetConfigmapsList() (*apiv1.ConfigMapList, error) {
	clientSet, err := client.ClientSet()
	if err != nil {
		return nil, err
	}
	config, err := clientSet.CoreV1().ConfigMaps("").List(context.Background(), metav1.ListOptions{})
	return config, err
}

func GetConfigmapsListByLabels(namespace, labels string) (*apiv1.ConfigMapList, error) {
	clientSet, err := client.ClientSet()
	if err != nil {
		return nil, err
	}
	listoptions := metav1.ListOptions{
		LabelSelector: labels,
	}

	config, err := clientSet.CoreV1().ConfigMaps(namespace).List(context.Background(), listoptions)
	return config, err
}

func GetConfigmapsListByOnlyLabels(labels string) (*apiv1.ConfigMapList, error) {
	clientSet, err := client.ClientSet()
	if err != nil {
		return nil, err
	}
	listoptions := metav1.ListOptions{
		LabelSelector: labels,
	}
	config, err := clientSet.CoreV1().ConfigMaps("").List(context.Background(), listoptions)
	return config, err
}

func GetConfigmapsByName(namespace, name string) (*apiv1.ConfigMap, error) {
	clientSet, err := client.ClientSet()
	if err != nil {
		return nil, err
	}
	config, err := clientSet.CoreV1().ConfigMaps(namespace).Get(context.Background(), name, metav1.GetOptions{})
	return config, err
}

func DeleteConfigmaps(namespace, name string) error {
	clientSet, err := client.ClientSet()
	if err != nil {
		return err
	}

	err = clientSet.CoreV1().ConfigMaps(namespace).Delete(context.Background(), name, metav1.DeleteOptions{})
	return err
}

// TODO : CreateConfigmap
//func CreateConfigmap(data models.ConfigmapsJson) (*apiv1.ConfigMap, error) {
//
//}

// TODO : UpdateConfigmap
//func UpdateConfigmap(data models.ConfigmapsJson) (*apiv1.ConfigMap, error) {
//
//}
