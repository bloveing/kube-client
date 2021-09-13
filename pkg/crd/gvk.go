package crd

import (
	"fmt"
	"github.com/bloveing/kube-client/pkg/client"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type GVKStruct struct {
	Group   string
	Version string
	Kind    string
	Name    string
}

func GVKList() (GVKList map[string]GVKStruct, err error) {
	discoveryClient, err := client.DiscoveryClient()
	if err != nil {
		panic(err)
	}
	_, APIResourceList, err := discoveryClient.ServerGroupsAndResources()
	if err != nil {
		panic(err)
	}
	for _, list := range APIResourceList {
		gv, err := schema.ParseGroupVersion(list.GroupVersion)
		if err != nil {
			panic(err)
		}
		for _, resource := range list.APIResources {
			fmt.Printf("name:%v,group:%v,version:%v\n", resource.Name, gv.Group, gv.Version)
			fmt.Println(resource.Name)

		}
	}
	return nil, nil

}
