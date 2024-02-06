package main

import (
	"flag"
	"github.com/zhangliweixyz/etcd-controller/controller"
	"time"

	"github.com/golang/glog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	clientset "github.com/zhangliweixyz/etcd-controller/pkg/generated/clientset/versioned"
	informers "github.com/zhangliweixyz/etcd-controller/pkg/generated/informers/externalversions"
)

var (
	masterURL  string
	kubeconfig string
)

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}

func main() {
	flag.Parse()

	// 处理入参
	kubeconfig = "/Users/admin/.kube/config"
	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %s", err.Error())
	}
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	etcdClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building example clientset: %s", err.Error())
	}

	etcdInformerFactory := informers.NewSharedInformerFactory(etcdClient, time.Second*30)
	//得到controller
	etcdController := controller.NewController(kubeClient, etcdClient,
		etcdInformerFactory.Etcd().V1beta2().EtcdClusters())

	//启动informer
	stopCh := make(chan struct{})
	go etcdInformerFactory.Start(stopCh)

	//controller开始处理消息
	if err = etcdController.Run(2, stopCh); err != nil {
		glog.Fatalf("Error running controller: %s", err.Error())
	}
}

// go run main.go -alsologtostderr=true
