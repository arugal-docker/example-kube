package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/arugal-docker/example-kube/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var log *logrus.Logger

func init() {
	log = logger.Log
}

var (
	config = Config{
		Addr:      ":8080",
		Namespace: apiv1.NamespaceDefault,
	}
	triggerCh = make(chan Trigger, 10)
	once      sync.Once
	clientSet *kubernetes.Clientset
)

func gotify(title string, message string) {
	if config.GotifyAddress == "" || config.GotifyToken == "" {
		return
	}

	client := resty.New()

	format := make(map[string]string)
	format["title"] = title
	format["message"] = message

	_, err := client.R().
		SetFormData(format).
		Post(fmt.Sprintf("http://%s/message?token=%s", config.GotifyAddress, config.GotifyToken))
	if err != nil {
		log.Errorf("gotify error, err:%s address:%s, token:%s", err, config.GotifyAddress, config.GotifyToken)
	}
}

func handler(namespace string, trigger Trigger) {
	once.Do(func() {
		config, err := clientcmd.BuildConfigFromFlags("", config.kubeConfig)
		if err != nil {
			panic(err)
		}
		clientSet, err = kubernetes.NewForConfig(config)
		if err != nil {
			panic(err)
		}
	})

	podsClient := clientSet.CoreV1().Pods(namespace)

	pods, err := podsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Errorf("%s pod list error, err: %v", namespace, err)
		return
	}
	for _, p := range pods.Items {
		for _, c := range p.Spec.Containers {
			if c.Image == fmt.Sprintf("registry.%s.aliyuncs.com/%s:%s", trigger.Repository.Region, trigger.Repository.RepoFullName, trigger.PushData.Tag) {
				err := podsClient.Delete(context.TODO(), p.Name, metav1.DeleteOptions{})
				if err != nil {
					log.Errorf("%s delete error, err: %v", p.Name, err)
				}
				gotify("pod update", fmt.Sprintf("delete pod %s/%s", p.Namespace, p.Name))
				log.Infof("delete pod %s/%s", p.Namespace, p.Name)
				break
			}
		}
	}
}

func main() {
	config.addFlags()
	flag.Parse()

	r := gin.Default()
	r.POST("/trigger", func(c *gin.Context) {
		if c.Request.Body != nil {
			defer c.Request.Body.Close()
			body, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				log.Errorf("Read body err: %v", err)
				return
			}
			log.Infof("received trigger: %s", string(body))
			trigger := Trigger{}
			err = json.Unmarshal(body, &trigger)
			if err != nil {
				log.Infof("Unmarshal err: %v", err)
				return
			}
			triggerCh <- trigger
		}
	})
	go func() {
		for trigger := range triggerCh {
			handler(config.Namespace, trigger)
		}
	}()
	_ = r.Run(config.Addr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
