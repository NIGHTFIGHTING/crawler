package main

import (
    "net/http"
    "io/ioutil"
    "fmt"
    "regexp"
)

func main() {
    resp, err := http.Get(
        "http://www.zhenai.com/zhenghun")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        fmt.Println("Error: status code", resp.StatusCode)
        return
    }
    // 1.utf8Reader := transform.NewReader(resp.Body,
    //         simplifiedchinese.GBK.NewDecoder())
    // 2.自动确认编码
    // e := determineEncoding(resp.Body)
    // utf8Reader := transform.NewReader(resp.Body,
    //         e.NewDecoder())
    all, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    //fmt.Printf("%s\n", all)
    printCityList(all)
}

func printCityList(contents []byte) {
    re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
    //matches := re.FindAll(contents, -1)
    //for _,m := range matches {
    //    fmt.Printf("%s\n", m)
    //}
    matches := re.FindAllSubmatch(contents, -1)
    for _, m := range matches {
//        for _, subMatch := range m {
//            fmt.Printf("%s", subMatch)
//        }
//        fmt.Println()
        fmt.Printf("City:%s, URL:%s\n", m[2], m[1])
    }
    fmt.Printf("Matches found:%d \n", len(matches))
}

//func determineEncoding(
//    r io.Reader) encoding.Encoding {
//    bytes, err := bufio.NewReader(r).Peek(1024)
//    if err != nil {
//        panic(err)
//    }
//    e, _, _ := charset.DetermineEncoding(
//        bytes, "") 
//    return e
//}
