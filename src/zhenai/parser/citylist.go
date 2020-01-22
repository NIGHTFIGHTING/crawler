package parser

import (
    "engine"
    "regexp"
    "fmt"
)

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
