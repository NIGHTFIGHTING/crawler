package parser

import (
    "testing"
    //. "zhenai/parser"
    "fetcher"
    "fmt"
    //"io/ioutil"
)

func TestParseCityList(t *testing.T) {
    contents, err := fetcher.Fetcher(
        "https://album.zhenai.com/u/1727435860")
    if err != nil {
        panic(err)
    }
    fmt.Printf("%s\n", contents)
//    contents, err := ioutil.ReadFile(
//        "profile_test_data.html")
//    if err != nil {
//        panic(err)
//    }
//    fmt.Printf("%s\n", contents)
//    result := ParseProfile(contents)
//    if len(result.Items) != 1 {
//        t.Errorf("Items should contain 1 "+
//            "element; but was %v", result.Items)
//    }
//
//    profile := result.Items[0].(model.Profile)
//
//    expected := model.Profile{
//
//    }
}
