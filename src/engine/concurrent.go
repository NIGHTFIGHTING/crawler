// 15-4队列实现调度器
package engine

import (
    "log"
)

type ConcurrentEngine struct {
    Scheduler Scheduler
    WorkerCount int
}

type Scheduler interface {
    Submit(Request)
    ConfigureMasterWorkerChan(chan Request)
    WorkerReady(chan Request)
    Run()
}
func (e *ConcurrentEngine) Run(seeds ...Request) {

    //in := make(chan Request)
    out := make(chan ParseResult)
    // schduler
    //e.Scheduler.ConfigureMasterWorkerChan(in)
    e.Scheduler.Run()

    for i := 0; i < e.WorkerCount; i++ {
        // 启动的worker协程从shcduler获得任务
        //createWorker(in, out)
        createWorker(out, e.Scheduler)
    }
    for _, r := range seeds {
        // engine向scheduler提交任务
        e.Scheduler.Submit(r)
    }

    itemCount := 0
    for {
        result := <- out
        for _, item := range result.Items {
            log.Printf("Got item #%d: %v", itemCount, item)
            itemCount++
        }
        for _, request := range result.Requests {
            e.Scheduler.Submit(request)
        }
    }
}

func createWorker(
    //in chan Request, out chan ParseResult) {
    out chan ParseResult, s Scheduler) {
    in := make(chan Request)
    go func() {
        for {
            s.WorkerReady(in)
            // 从scheduler获得任务
            request := <- in
            result, err := worker(request)
            if err != nil {
                continue
            }
            out <- result
        }
    }()
}
