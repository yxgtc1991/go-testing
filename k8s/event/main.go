package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const IsCanary = true

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	kubeConfig, err := clientcmd.BuildConfigFromFlags("", dir+"/k8s/kubeconfig")
	if err != nil {
		panic(err)
	}

	clientSet, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		panic(err)
	}

	sharedFactory := informers.NewSharedInformerFactoryWithOptions(clientSet, time.Minute, informers.WithTweakListOptions(func(options *metav1.ListOptions) {
		options.LabelSelector = getCanaryLabelSelector(getCanaryUser(), map[string]string{}).String()
	}))

	sharedFactory.Core().V1().Events().Informer()

	stopCh := make(chan struct{})
	sharedFactory.Start(stopCh)

	for v, synced := range sharedFactory.WaitForCacheSync(stopCh) {
		if !synced {
			fmt.Printf("sync %v fail", v)
		}
	}

	ns := "9d302cfb"
	events, err := sharedFactory.Core().V1().Events().Lister().Events(ns).List(labels.NewSelector())
	if err != nil {
		panic(err)
	}
	for i, event := range events {
		succReason := "forceMaster switch success"
		if event.Reason == succReason {
			fmt.Println("id: ", i)
			fmt.Println("ns: ", event.Namespace)
			fmt.Println("app: ", event.InvolvedObject.Name)
			fmt.Println("tenant: ", event.Labels["TenantId"])
			fmt.Printf("msg: %s\n\n", event.Message)
		}
		failReason := "forceMaster switch failed"
		if event.Reason == failReason {
			fmt.Println("id: ", i)
			fmt.Println("ns: ", event.Namespace)
			fmt.Println("app: ", event.InvolvedObject.Name)
			fmt.Println("tenant: ", event.Labels["TenantId"])
			fmt.Printf("msg: %s\n\n", event.Message)
		}
	}
}

func getCanaryLabelSelector(canaryUsers []string, labelMap map[string]string) labels.Selector {
	if canaryUsers == nil {
		canaryUsers = []string{""}
	}

	var requirement *labels.Requirement
	if IsCanary {
		requirement, _ = labels.NewRequirement("TenantId", selection.In, canaryUsers)
	} else {
		requirement, _ = labels.NewRequirement("TenantId", selection.NotIn, canaryUsers)
	}

	selector := labels.NewSelector().Add(*requirement)
	if labelMap != nil {
		for k, v := range labelMap {
			requirement2, _ := labels.NewRequirement(k, selection.Equals, []string{v})
			selector = selector.Add(*requirement2)
		}
	}
	return selector
}

func getCanaryUser() []string {
	users := strings.Split("CIDC-U-ed0608fc0a2d4f8ca32f80e49d302cfb", "|")
	return users
}
