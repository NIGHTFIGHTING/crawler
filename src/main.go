package main

import (
    "zhenai/parser"
    "engine"
)

func main() {
    engine.Run(engine.Request{
        Url: "http://www.zhenai.com/zhenghun",
        ParserFunc: parser.ParseCityList,
    })
}
