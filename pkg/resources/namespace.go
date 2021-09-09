package resources

import (
	"context"
	"github.com/bloveing/kube-client/pkg/client"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNamespacesByName(name string) (*apiv1.Namespace, error) {
	clientSet, err := client.ClientSet()
	if err != nil {
		return nil, err
	}
	namespace, err := clientSet.CoreV1().Namespaces().Get(context.Background(), name, metav1.GetOptions{})
	return namespace, err
}

func GetNamespaces() (*apiv1.NamespaceList, error) {
	clientSet, err := client.ClientSet()
	if err != nil {
		return nil, err
	}
	namespaces, err := clientSet.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	return namespaces, err
}

func CreateNamespaces(name string) error {
	clientSet, err := client.ClientSet()
	if err != nil {
		return err
	}
	ns := &apiv1.Namespace{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	_, err = clientSet.CoreV1().Namespaces().Create(context.Background(), ns, metav1.CreateOptions{})
	return err
}

func DeleteNamespaces(name string) error {
	clientSet, err := client.ClientSet()
	if err != nil {
		return err
	}
	err = clientSet.CoreV1().Namespaces().Delete(context.Background(), name, metav1.DeleteOptions{})
	return err
}
