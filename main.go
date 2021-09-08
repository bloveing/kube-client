package main

import (
	"fmt"
	"github.com/bloveing/kube-client/pkg/client"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func main() {
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
            fmt.Printf("name:%v, \t\t group:%v, \t\t version:%v\n", resource.Name, gv.Group, gv.Version)
        }
    }

}
