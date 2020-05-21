module github.com/arugal-docker/example-kube

go 1.13

require (
	github.com/gin-gonic/gin v1.5.0
	github.com/go-resty/resty/v2 v2.2.0
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.5
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	istio.io/client-go v0.0.0-20200505182340-146ba01d5357
	k8s.io/api v0.18.1
	k8s.io/apimachinery v0.18.1
	k8s.io/client-go v0.18.1
)
