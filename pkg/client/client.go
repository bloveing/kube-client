package client

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	scheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	if h := os.Getenv("USERPROFILE"); h != "" { // windows
		return h
	}

	return "/root"
}

func InitConfig() (*rest.Config, error) {
	//var kubeconfig *string
	//if home := homeDir(); home != "" {
	//	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	//} else {
	//	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	//}
	//flag.Parse()
	path, _ := os.Getwd()
	kubeConfig := filepath.Join(path, ".kube", "config")

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		return nil, err
	}
	return config, err

}

func RESTClientCore() (*rest.RESTClient, error) {
	var once sync.Once
	var restClient *rest.RESTClient
	var err error
	once.Do(func() {
		log.Print("start InitConfig()")
		config, _ := InitConfig()

		///设置config.APIPath请求的HTTP路径
		config.APIPath = "api"
		//设置config.GroupVersion请求的资源组/资源版本
		config.GroupVersion = &corev1.SchemeGroupVersion
		//设置config.NegotiatedSerializer数据的解码器
		config.NegotiatedSerializer = scheme.Codecs

		//实例化Client对象
		restClient, err = rest.RESTClientFor(config)

	})
	return restClient, nil
}

// use sync.Once, so clientset can only init once
func ClientSet() (*kubernetes.Clientset, error) {
	var once sync.Once
	var clientset *kubernetes.Clientset
	var err error
	once.Do(func() {
		log.Print("start InitConfig()")
		config, _ := InitConfig()
		// create the clientset
		clientset, err = kubernetes.NewForConfig(config)
	})
	return clientset, nil
}

func DynamicClient() (dynamic.Interface, error) {
	var once sync.Once
	var dynamicClient dynamic.Interface
	var err error

	once.Do(func() {
		log.Print("start InitConfig()")
		config, _ := InitConfig()
		// create the clientset
		dynamicClient, err = dynamic.NewForConfig(config)

	})
	return dynamicClient, nil
}

func DiscoveryClient() (*discovery.DiscoveryClient, error) {
	var once sync.Once
	var discoveryClient *discovery.DiscoveryClient
	var err error

	once.Do(func() {
		log.Print("start InitConfig()")
		config, _ := InitConfig()
		// create the clientset
		discoveryClient, err = discovery.NewDiscoveryClientForConfig(config)
	})
	return discoveryClient, nil
}
