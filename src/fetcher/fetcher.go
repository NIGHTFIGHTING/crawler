package fetcher

import (
    "net/http"
    "io/ioutil"
    "fmt"
    "regexp"
    //"golang.org./x/net/html/chatset"
    //"bufio"
    //"log"
)

func Fetcher(url string)([]byte, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        fmt.Println("Error: status code", resp.StatusCode)
        //return nil, erros.New()
        return nil, fmt.Errorof ("Wrong status code: %d",
            resp.StatusCode)
    }
    // 自动确认编码
    // e := determineEncoding(resp.Body)
    // utf8Reader := transform.NewReader(resp.Body,
    //         e.NewDecoder())
    //return ioutil.ReadAll(utf8Reader)
    return ioutil.ReadAll(resp.Body)
}

//func determineEncoding(
//    r io.Reader) encoding.Encoding {
//    bytes, err := bufio.NewReader(r).Peek(1024)
//    if err != nil {
//        log.Printf("Fetcher error: %v", err)
//        return unicode.UTF8
//    }
//    e, _, _ := charset.DetermineEncoding(
//        bytes, "") 
//    return e
//}
