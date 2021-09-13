### model - v1

```shell
go get -u k8s.io/client-go@v0.22.1  
go get github.com/elastic/cloud-on-k8s
```


---
### 介绍

> 没有抽取公共的create / apply 等方法 , 针对每个资源做的 SDK
> 
>> V2版本 实现接口统一


### 参考
[git/lflxp-kubectl](https://github.com/lflxp/lflxp-kubectl)   
[git/client-go-helper](https://github.com/ica10888/client-go-helper/blob/master/pkg/kubectl/create.go)  
[git/kubectl-warp](https://github.com/ernoaapa/kubectl-warp/blob/master/pkg/kubectl/client.go)

[git/kubectl](https://github.com/kubernetes/kubectl)  
[git/klient](https://github.com/johandry/klient)  
[blog/kclient-paper](http://blog.johandry.com/post/build-k8s-client/)  


### ECK
eck 的model:/pkg/apis/elasticsearch/v1/elasticsearch_types.go
eck 的model：\pkg\apis\elasticsearch\v1\elasticsearch_types.go
