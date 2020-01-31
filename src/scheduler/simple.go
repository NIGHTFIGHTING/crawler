package scheduler

import (
    "engine"
)

type SimpleScheduler struct {
    workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(
    r engine.Request) {
    // send request down to worker chan
    go func() {s.workerChan <- r}()
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(
    c chan engine.Request){
    s.workerChan = c
}

func (s *SimpleScheduler) WorkerReady(
    w chan engine.Request) {
}
func (s *SimpleScheduler) Run() {
}
