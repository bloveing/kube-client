package resources

import (
	"context"
	"github.com/bloveing/kube-client/pkg/client"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetEvenstList() (*apiv1.EventList, error) {
	clientSet, err := client.ClientSet()
	if err != nil {
		return nil, err
	}

	config, err := clientSet.CoreV1().Events("").List(context.Background(), metav1.ListOptions{})
	return config, err
}

func GetEventsByNamespace(namespace string) (*apiv1.EventList, error) {
	clientSet, err := client.ClientSet()
	if err != nil {
		return nil, err
	}

	config, err := clientSet.CoreV1().Events(namespace).List(context.Background(), metav1.ListOptions{})
	return config, err
}

func GetEventsByName(namespace, name string) (*apiv1.Event, error) {
	clientSet, err := client.ClientSet()
	if err != nil {
		return nil, err
	}

	config, err := clientSet.CoreV1().Events(namespace).Get(context.Background(), name, metav1.GetOptions{})
	return config, err
}

func DeleteEvents(namespace, name string) error {
	clientSet, err := client.ClientSet()
	if err != nil {
		return err
	}

	err = clientSet.CoreV1().Events(namespace).Delete(context.Background(), name, metav1.DeleteOptions{})
	return err
}
