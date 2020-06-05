package main
import (
    "fmt"
    "log"
    "k8s.io/client-go/tools/clientcmd"
    "path/filepath"
    v1 "k8s.io/api/core/v1"
    "k8s.io/client-go/kubernetes"
    "os"
//    "context"
//   apiv1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

)

// proper example at: https://github.com/vladimirvivien/k8s-client-examples/blob/master/go/pvcwatch/main.go
func main() {
    kubeconfig := filepath.Join(
         os.Getenv("HOME"), ".kube", "config",
    )
    config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
    if err != nil {
        log.Fatal(err)
    }
    clientset, err := kubernetes.NewForConfig(config)  // erializer client to let us access API objects. Type Clientset, from package kubernetes
    if err != nil {
        log.Fatal(err)
    }
    api := clientset.CoreV1() // check which one we need to use from struct defination of clientset here https://github.com/kubernetes/client-go/blob/master/kubernetes/clientset.go#L113
    listOptions := metav1.ListOptions{
    }
    pvcs, err := api.PersistentVolumeClaims("default").List(listOptions) // default namsepace
    watcher, err :=api.PersistentVolumeClaims("default").Watch(listOptions)
    ch := watcher.ResultChan()
    for event := range ch {
        pvc, OK := event.Object.(*v1.PersistentVolumeClaim)
//        fmt.Println(pvc)
        fmt.Println(pvc.Spec)
        x:=pvc.Spec.Resources.Requests["storage"]
        fmt.Println("\n",x.String())
        fmt.Println("\n",OK,"\n", event.Type)
    }





    pods, err := api.Pods("default").List(listOptions) //  more functions with api(Corev1) https://godoc.org/k8s.io/client-go/kubernetes/typed/core/v1
    for _, pod := range pods.Items {
    fmt.Println("\n",pod.Name)
    }
    printPVCs(pvcs)
}



func printPVCs(pvcs *v1.PersistentVolumeClaimList) {
    template := "%-32s%-8s%-8s\n"
    fmt.Printf(template, "NAME", "STATUS", "CAPACITY")
    for _, pvc := range pvcs.Items {
        quant := pvc.Spec.Resources.Requests[v1.ResourceStorage]
        fmt.Printf(
            template,
            pvc.Name,
            string(pvc.Status.Phase),
            quant.String())
    }
}
