package main

import (
    "zhenai/parser"
    "engine"
    "scheduler"
)

func main() {
    e := engine.ConcurrentEngine{
        Scheduler: &scheduler.QueueScheduler{},
        WorkerCount: 10,
    }
    e.Run(engine.Request{
        Url: "http://www.zhenai.com/zhenghun",
        ParserFunc: parser.ParseCityList,
    })
//    1.单机版本
//    engine.SimpleEngine{}.Run(engine.Request{
//        Url: "http://www.zhenai.com/zhenghun",
//        ParserFunc: parser.ParseCityList,
//    })
//
//    2.simple.go和并发版本调度器
//    e := engine.ConcurrentEngine{
//        Scheduler: &scheduler.SimpleScheduler{},
//        WorkerCount: 10,
//    }
}
