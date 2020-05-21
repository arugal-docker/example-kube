package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/arugal-docker/example-kube/logger"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/metadata"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var log *logrus.Logger

func init() {
	log = logger.Log
}

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	namespace := *flag.String("namespace", "kube-system", "")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	metadataClient, err := metadata.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	deploymentsClient := clientset.AppsV1().Deployments(namespace)

	// List Deployments
	prompt()
	fmt.Printf("Listing deployments in namespace %q:\n", namespace)
	list, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, d := range list.Items {
		fmt.Printf(" * %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
	}

	podsClient := clientset.CoreV1().Pods(namespace)

	// List Pods
	prompt()
	pods, err := podsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, p := range pods.Items {
		nodeResource := schema.GroupVersionResource{Group: "", Version: "v1", Resource: "nodes"}
		raw, err := metadataClient.Resource(nodeResource).Get(context.TODO(), p.Spec.NodeName, metav1.GetOptions{})
		if err != nil {
			log.Warnf("unable to get node %q for pod %q: %v", p.Spec.NodeName, p.Name, err)
			return
		}
		nodeMeta, err := meta.Accessor(raw)
		if err != nil {
			log.Warnf("unable to get node meta: %v", nodeMeta)
			return
		}
		log.Println(nodeMeta)
		fmt.Printf(" pod name %s \n", p.Name)
	}
}

func prompt() {
	fmt.Printf("-> Press Return key to continue.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println()
}
