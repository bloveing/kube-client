package main

import (
	"context"
	"fmt"
	"github.com/bloveing/kube-client/pkg/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	clientSet, err := client.ClientSet()
	if err != nil {
		fmt.Println(err)
	}
	//d, err := clientSet.AppsV1().Deployments("").List(context.Background(),metav1.ListOptions{})
	d, err := clientSet.AppsV1().Deployments("kubeapps").List(context.Background(), metav1.ListOptions{})
	fmt.Println(d)
	clientSet2, err := client.ClientSet()
	d2, err := clientSet2.AppsV1().Deployments("k8s-api").List(context.Background(), metav1.ListOptions{})
	fmt.Println(d2)
}
