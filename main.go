package main

import (
	"flag"
	"fmt"
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeConfig = flag.String("kubeconfig", "/home/bisakh/.kube/config", "kube config file")

func main() {
	flag.Parse()

	// Outside Cluster
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		fmt.Println("Unable to get config file: ", err)

		// In cluster config method
		config, err = rest.InClusterConfig()
		if err != nil {
			log.Panicln("Application is neither running inside cluster too: ", err)
		}
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Panicln("Unable to get client sets: ", err)
	}

	// Fetch server version
	ver, err := clientSet.ServerVersion()
	if err != nil {
		log.Panicln("Failed to fetch server version: ", err)
	}
	fmt.Printf("Kubernetes Version: %s.%s\n", ver.Major, ver.Minor)
}
