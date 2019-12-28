package  timer 

import(
	"time"
	"log"
	"fmt"
	"context"

	"github.com/robfig/cron"
	batchv1beta1 "github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type (
	TimeService struct {
		logger *log.Logger
		kubeclientset client.Client
		triggers map[string]*timerTriggerWithCron
	}

	timerTriggerWithCron struct {
		trigger batchv1beta1.TimerTrigger
		cron    *cron.Cron
	}
)
func MakeTimeService(logger *log.Logger, clientset client.Client) *TimeService {
	ts := &TimeService{
		logger: logger,
		kubeclientset: clientset,
		triggers: make(map[string]*timerTriggerWithCron),
	}
	go ts.svc()
	
	return ts
}

func (ts *TimeService) svc() {
	for{
		timerList := &batchv1beta1.TimerTriggerList{}
		err := ts.kubeclientset.List(context.TODO(), timerList)
		if err != nil {
			ts.logger.Fatal(err)
		}
		ts.syncCron(timerList)
		
		time.Sleep(3*time.Second)
	}
}

func (ts *TimeService) syncCron( triggers *batchv1beta1.TimerTriggerList){
	triggerMap := make(map[string]bool)
	for _, t := range triggers.Items {
		triggerMap[cacheKey(&t.ObjectMeta)] = true
		if item, found := ts.triggers[cacheKey(&t.ObjectMeta)]; found{
			if item.trigger.Spec.Cron != t.Spec.Cron{
				if item.cron != nil {
					item.cron.Stop()
					ts.logger.Println("Stop cron for time trigger:", "trigger-", t.ObjectMeta.Name)
				}
				item.cron = ts.newCron(t)
			}
			item.trigger = t
		} else {
			ts.triggers[cacheKey(&t.ObjectMeta)] = &timerTriggerWithCron{
				trigger: t,
				cron: ts.newCron(t),
			}
		}
	}

	for k, v := range ts.triggers {
		if _, found := triggerMap[k]; !found{
			if v.cron != nil {
				v.cron.Stop()
				ts.logger.Println("Remove cron for time trigger:", "trigger-", v.trigger.ObjectMeta.Name)
			}
			delete(ts.triggers, k)
		}
	}
}

func (ts *TimeService) newCron(t batchv1beta1.TimerTrigger) *cron.Cron{
	c := cron.New() 
	ts.logger.Println("New Cron:",t.Spec.Cron)
	c.AddFunc(t.Spec.Cron, func() {
		fmt.Println("test test")
		ts.logger.Println(fmt.Sprintf("Got the timer:%v with fuction: %v",t.ObjectMeta.Name, t.Spec.JobReference))
		// return
	})
	c.Start()
	ts.logger.Println("added new cron for time trigger:", "trigger-", t.ObjectMeta.Name)
	return c
}