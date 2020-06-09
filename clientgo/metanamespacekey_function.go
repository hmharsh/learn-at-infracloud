package main

import (
        "fmt"
        corev1 "k8s.io/api/core/v1"
        "k8s.io/apimachinery/pkg/util/runtime"
        "log"
        "os"
        "path/filepath"

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
        informer := factory.Core().V1().Events().Informer()
        stopper := make(chan struct{})
        defer close(stopper)
        defer runtime.HandleCrash()
        informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
                AddFunc: onAdd,
        })
        go informer.Run(stopper)
        if !cache.WaitForCacheSync(stopper, informer.HasSynced) {
                runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
                return
        }
        <-stopper

}

func onAdd(obj interface{}) {
        str, _ := cache.MetaNamespaceKeyFunc(obj)
        fmt.Println(str)
        event := obj.(*corev1.Event)
        fmt.Println(event.InvolvedObject.Kind, "kind of event fired of type", event.Type)
}
