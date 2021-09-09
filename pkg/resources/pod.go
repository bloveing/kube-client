package resources

import (
	"bytes"
	"context"
	"github.com/bloveing/kube-client/pkg/client"
	"io"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetPodList() (*apiv1.PodList, error) {
	clientSet, err := client.ClientSet()
	if err != nil {
		return nil, err
	}

	config, err := clientSet.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	return config, err
}

func GetPodListByNamespace(namespace string) (*apiv1.PodList, error) {
	clientSet, err := client.ClientSet()
	if err != nil {
		return nil, err
	}

	config, err := clientSet.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	return config, err
}

func GetPodListByLabels(namespace, label string) (*apiv1.PodList, error) {
	clientSet, err := client.ClientSet()
	if err != nil {
		return nil, err
	}

	listoptions := metav1.ListOptions{
		LabelSelector: label,
	}

	config, err := clientSet.CoreV1().Pods(namespace).List(context.Background(), listoptions)
	return config, err
}

func GetPodByName(namespace, name string) (*apiv1.Pod, error) {
	clientSet, err := client.ClientSet()
	if err != nil {
		return nil, err
	}

	config, err := clientSet.CoreV1().Pods(namespace).Get(context.Background(), name, metav1.GetOptions{})
	return config, err
}

func DeletePod(namespace, name string) error {
	clientSet, err := client.ClientSet()
	if err != nil {
		return err
	}

	err = clientSet.CoreV1().Pods(namespace).Delete(context.Background(), name, metav1.DeleteOptions{})
	return err
}

func GetPodLogByPodId(namespace, podid string) (string, error) {
	var result string
	f := func(s int64) *int64 {
		return &s
	}

	clientSet, err := client.ClientSet()
	if err != nil {
		return result, err
	}

	options := apiv1.PodLogOptions{
		Follow:    false,
		TailLines: f(1000),
	}
	req := clientSet.CoreV1().Pods(namespace).GetLogs(podid, &options)
	readCloser, err := req.Stream(context.Background())
	if err != nil {
		return result, err
	}
	defer readCloser.Close()
	var out bytes.Buffer
	_, err = io.Copy(&out, readCloser)
	if err != nil {
		return result, err
	}
	result = out.String()
	return result, nil
}

func GetPodLogByPodIdByNum(namespace, podid string, num int64) (string, error) {
	var result string
	f := func(s int64) *int64 {
		return &s
	}

	clientSet, err := client.ClientSet()
	if err != nil {
		return result, err
	}

	options := apiv1.PodLogOptions{
		Follow:       false,
		SinceSeconds: f(num),
	}
	req := clientSet.CoreV1().Pods(namespace).GetLogs(podid, &options)
	readCloser, err := req.Stream(context.Background())
	if err != nil {
		return result, err
	}
	defer readCloser.Close()
	var out bytes.Buffer
	_, err = io.Copy(&out, readCloser)
	if err != nil {
		return result, err
	}
	result = out.String()
	return result, nil
}

func GetPodLogByPodIdAll(namespace, podid string) (string, error) {
	var result string

	clientSet, err := client.ClientSet()
	if err != nil {
		return result, err
	}

	options := apiv1.PodLogOptions{
		Follow: false,
	}
	req := clientSet.CoreV1().Pods(namespace).GetLogs(podid, &options)
	readCloser, err := req.Stream(context.Background())
	if err != nil {
		return result, err
	}
	defer readCloser.Close()
	var out bytes.Buffer
	_, err = io.Copy(&out, readCloser)
	if err != nil {
		return result, err
	}
	result = out.String()
	return result, nil
}
