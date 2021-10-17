package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func ListPods(ctx context.Context, cl *kubernetes.Clientset, ns string) ([]string, error) {
	pods, err := cl.CoreV1().Pods(ns).List(ctx, metav1.ListOptions{
		//LabelSelector: "hello-world",
	})
	if err != nil {
		return nil, err
	}

	var names []string
	for _, pod := range pods.Items {
		names = append(names, pod.Name)
	}

	return names, nil
}

func ListDeployments(ctx context.Context, cl *kubernetes.Clientset, ns string) (names []string, err error) {
	deps, err := cl.AppsV1().Deployments(ns).List(ctx, metav1.ListOptions{})

	if err != nil {
		return
	}
	for _, d := range deps.Items {
		names = append(names, d.Name)
	}
	return
}

var kubeConfig = flag.String("kubeconfig", "/home/bisakh/.kube/config", "kube config file")

func main() {
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		log.Fatalln("Unable to get config: ", err)
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Panicln("Unable to get client sets: ", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// List Pods

	podNames, err := ListPods(ctx, clientSet, "dev")
	if err != nil {
		log.Println("Unable to get pods: ", err)
	}
	fmt.Println("PodNames: ", podNames)

	// List Deployments

	depNames, err := ListDeployments(ctx, clientSet, "dev")
	if err != nil {
		log.Println("Unable to get pods: ", err)
	}
	fmt.Println("Deployment Names: ", depNames)
}
