package main

import (
        "fmt"
        "log"
        //    "k8s.io/client-go/tools/clientcmd"
        "path/filepath"
        //    apiv1 "k8s.io/api/core/v1"
        //    "k8s.io/client-go/kubernetes"
        "os"
        //      "k8s.io/apimachinery/pkg/util/runtime"
        //    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
        //    "k8s.io/client-go/informers"
        //    "k8s.io/client-go/tools/cache"
        //

        corev1 "k8s.io/api/core/v1"
        "k8s.io/apimachinery/pkg/util/runtime"

        "k8s.io/client-go/informers"
        "k8s.io/client-go/kubernetes"
        "k8s.io/client-go/tools/cache"
        "k8s.io/client-go/tools/clientcmd"
)

func main() {
        kubeconfig := filepath.Join(
                os.Getenv("HOME"), ".kube", "config",
        )
        config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
        if err != nil {
                log.Fatal(err)
        }
        clientset, err := kubernetes.NewForConfig(config)
        if err != nil {
                log.Fatal(err)
        }
        factory := informers.NewSharedInformerFactory(clientset, 0)
        informer := factory.Core().V1().Pods().Informer()
        stopper := make(chan struct{})
        defer close(stopper)
        defer runtime.HandleCrash()
        informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
                AddFunc: onAdd,
        })
        go informer.Run(stopper)
        // if !cache.WaitForCacheSync(stopper, informer.HasSynced) {
        //     runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
        //     return
        // }
        <-stopper

}

func onAdd(obj interface{}) {
        pod := obj.(*corev1.Pod)
        fmt.Println(pod.Name, "pod is created")
}
