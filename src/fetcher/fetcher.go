package fetcher

import (
    "net/http"
    "io/ioutil"
    "fmt"
    //"golang.org./x/net/html/chatset"
    //"bufio"
    //"log"
)

func Fetcher(url string)([]byte, error) {
//    resp, err := http.Get(url)
//    if err != nil {
//        return nil, err
//    }
    // 爬取网页数据时403
    // https://blog.csdn.net/qq_15977699/article/details/87932203
    // https://blog.csdn.net/qq_36183935/article/details/80499183
    client := http.Client{}
    request, err := http.NewRequest("GET", url, nil)
    if err!=nil {
        fmt.Printf("wrong http request: %s", err.Error())
        return nil, fmt.Errorf("wrong http request: %s", err.Error())
    }
    request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
    resp, err := client.Do(request)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Println("Error: status code", resp.StatusCode)
        //return nil, erros.New()
        return nil, fmt.Errorf ("Wrong status code: %d",
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
