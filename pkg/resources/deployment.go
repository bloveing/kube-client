package resources

import (
	"context"
	"github.com/bloveing/kube-client/pkg/client"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetDeploymentList() (*v1.DeploymentList, error) {
	clientSet, err := client.ClientSet()
	if err != nil {
		return nil, err
	}
	deploy, err := clientSet.AppsV1().Deployments("").List(context.Background(), metav1.ListOptions{})
	return deploy, err
}

func GetDeploymentListByNamespace(namespace string) (*v1.DeploymentList, error) {
	clientSet, err := client.ClientSet()
	if err != nil {
		return nil, err
	}
	deploy, err := clientSet.AppsV1().Deployments(namespace).List(context.Background(), metav1.ListOptions{})
	return deploy, err
}

func GetDeploymentByName(namespace, name string) (*v1.Deployment, error) {
	clientSet, err := client.ClientSet()
	if err != nil {
		return nil, err
	}
	deploy, err := clientSet.AppsV1().Deployments(namespace).Get(context.Background(), name, metav1.GetOptions{})
	return deploy, err
}
// TODO : create Deployment

func DeleteDeployment(namespace, name string) error {
	clientSet, err := client.ClientSet()
	if err != nil {
		return err
	}

	err = clientSet.AppsV1().Deployments(namespace).Delete(context.Background(), name, metav1.DeleteOptions{})
	return err
}

func ScaleDeployment(namespace, name string, num int32) (*v1.Deployment, error) {
	clientSet, err := client.ClientSet()
	if err != nil {
		return nil, err
	}
	deploy, err := GetDeploymentByName(namespace, name)
	if err != nil {
		return nil, err
	}

	deploy.Spec.Replicas = &num
	dps, err := clientSet.AppsV1().Deployments(namespace).Update(context.Background(), deploy, metav1.UpdateOptions{})
	return dps, err
}

func ChangeImageDeployment(namespace, name, image string) (*v1.Deployment, error) {
	clientSet, err := client.ClientSet()
	if err != nil {
		return nil, err
	}
	deploy, err := GetDeploymentByName(namespace, name)
	if err != nil {
		return nil, err
	}
	deploy.Spec.Template.Spec.Containers[0].Image = image
	dps, err := clientSet.AppsV1().Deployments(namespace).Update(context.Background(), deploy, metav1.UpdateOptions{})
	return dps, err
}
