package timer

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"github.com/robfig/cron"
	batchv1beta1 "github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type (
	TimeService struct {
		logger        logr.Logger
		kubeclientset client.Client
		routerUrl     string
		triggers      map[string]*timerTriggerWithCron
	}

	timerTriggerWithCron struct {
		trigger batchv1beta1.TimerTrigger
		cron    *cron.Cron
	}
)

func Start(in_logger logr.Logger, clientset client.Client, routerUrl string) {
	logger := in_logger.WithName("TimeService")
	ts := MakeTimeService(logger, clientset, routerUrl)
	ts.svc()
}

func MakeTimeService(logger logr.Logger, clientset client.Client, routerUrl string) *TimeService {
	timeLogger := logger.WithName("Timer")
	return &TimeService{
		logger:        timeLogger,
		kubeclientset: clientset,
		triggers:      make(map[string]*timerTriggerWithCron),
		routerUrl:     routerUrl,
	}
}

func (ts *TimeService) svc() {
	for {
		timerList := &batchv1beta1.TimerTriggerList{}
		err := ts.kubeclientset.List(context.TODO(), timerList)
		if err != nil {
			ts.logger.Error(err, "Failed to get Timer triggers")
			// os.Exit(1)
		}
		ts.syncCron(timerList)

		time.Sleep(3 * time.Second)
	}
}

func (ts *TimeService) syncCron(triggers *batchv1beta1.TimerTriggerList) {
	triggerMap := make(map[string]bool)
	for _, t := range triggers.Items {
		triggerMap[cacheKey(&t.ObjectMeta)] = true
		if item, found := ts.triggers[cacheKey(&t.ObjectMeta)]; found {
			if item.trigger.Spec.Cron != t.Spec.Cron {
				if item.cron != nil {
					item.cron.Stop()
					ts.logger.Info("Stop cron for time trigger ", "trigger-", t.ObjectMeta.Name)
				}
				item.cron = ts.newCron(t)
			}
			item.trigger = t
		} else {
			ts.triggers[cacheKey(&t.ObjectMeta)] = &timerTriggerWithCron{
				trigger: t,
				cron:    ts.newCron(t),
			}
		}
	}

	for k, v := range ts.triggers {
		if _, found := triggerMap[k]; !found {
			if v.cron != nil {
				v.cron.Stop()
				ts.logger.Info("Remove cron for time trigger:trigger ", "trigger-", v.trigger.ObjectMeta.Name)
			}
			delete(ts.triggers, k)
		}
	}
}

func (ts *TimeService) newCron(t batchv1beta1.TimerTrigger) *cron.Cron {
	c := cron.New()
	ts.logger.Info("New", "Cron", t.Spec.Cron)
	c.AddFunc(t.Spec.Cron, func() {
		// fmt.Println("test test")
		ts.logger.Info(fmt.Sprintf("Got the timer:%v with fuction: %v", t.ObjectMeta.Name, t.Spec.JobReference))
		makeHTTPRequest(ts.routerUrl + "/v1/" + t.Spec.JobReference.Name + "?namespace=" + t.Namespace)
		// return
	})
	c.Start()
	ts.logger.Info("added new cron for time trigger: ", "trigger-", t.ObjectMeta.Name)
	return c
}
