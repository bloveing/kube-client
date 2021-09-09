package resources

import (
	"context"
	"github.com/bloveing/kube-client/pkg/client"
	apiv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetJobList(namespace string) (*apiv1.JobList, error) {
	clientSet, err := client.ClientSet()
	if err != nil {
		return nil, err
	}

	config, err := clientSet.BatchV1().Jobs(namespace).List(context.Background(), metav1.ListOptions{})
	return config, err
}

func GetJobListByLabels(namespace, label string) (*apiv1.JobList, error) {
	clientSet, err := client.ClientSet()
	if err != nil {
		return nil, err
	}

	listoptions := metav1.ListOptions{
		LabelSelector: label,
	}

	config, err := clientSet.BatchV1().Jobs(namespace).List(context.Background(), listoptions)
	return config, err
}

func GetJobByName(namespace, name string) (*apiv1.Job, error) {
	clientSet, err := client.ClientSet()
	if err != nil {
		return nil, err
	}

	config, err := clientSet.BatchV1().Jobs(namespace).Get(context.Background(), name, metav1.GetOptions{})
	return config, err
}

func DeleteJob(namespace, name string) error {
	clientSet, err := client.ClientSet()
	if err != nil {
		return err
	}

	err = clientSet.BatchV1().Jobs(namespace).Delete(context.Background(), name, metav1.DeleteOptions{})
	return err
}
