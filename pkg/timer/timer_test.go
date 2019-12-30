package timer

import (
		"log"
		"os"
        // "time"
        "testing"

        batchv1beta1 "github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1"
        metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
        "k8s.io/apimachinery/pkg/runtime"
        "k8s.io/client-go/kubernetes/scheme"
        // "sigs.k8s.io/controller-runtime/pkg/client"
        "sigs.k8s.io/controller-runtime/pkg/client/fake"
) 

func TestMakeTimeService(t *testing.T) {
    timeTrigger := &batchv1beta1.TimerTrigger{
        ObjectMeta: metav1.ObjectMeta{
            Name: "timeTrigger",
            Namespace: "default",
        },
        Spec: batchv1beta1.TimerTriggerSpec{
            Cron: "0 */1 * * * *",
            JobReference: batchv1beta1.JobReference{
                Name: "example-httptrigger",
            },
        },
    }

    batchv1beta1.SchemeBuilder.AddToScheme(scheme.Scheme)
    
    // Objects to track in the fake client.
    objs := []runtime.Object{timeTrigger}

    // Create a fake client to mock API calls.
    cl := fake.NewFakeClient(objs...)
   
    var logger = log.New(os.Stdout,"", log.LstdFlags|log.Llongfile)
    MakeTimeService(logger,cl,"localhost:8000")
    // time.Sleep(30 * time.Second)
    for{

    }
}