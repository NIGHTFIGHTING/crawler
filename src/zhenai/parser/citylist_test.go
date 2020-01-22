package parser

import (
    "testing"
    . "zhenai/parser"
    //"fetcher"
    //"fmt"
    "io/ioutil"
)

func TestParseCityList(t *testing.T) {
//    contents, err := fetcher.Fetcher(
//        "http://www.zhenai.com/zhenghun")
//    if err != nil {
//        panic(err)
//    }
//    fmt.Printf("%s\n", contents)
    contents, err := ioutil.ReadFile(
        "citylist_test_data.html")
    if err != nil {
        panic(err)
    }
    //fmt.Printf("%s\n", contents)
    result := ParseCityList(contents)
    const resultSize = 470
    if len(result.Requests) != resultSize {
        t.Errorf("result should have %d" +
            "reuqests; but had %d",
            resultSize, len(result.Requests))
    }
}


