package main

import (
    "regexp"
    "fmt"
)

const text = `My email is flliuqi@didiglobal.com
email is abc@def.org
email2 is kkk@qq.com
email3 is ddd@abc.com.cn`

func main() {
    //re := regexp.MustCompile(".+@.+\\..+")
    re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)(\.[a-zA-Z0-9]+)`)
    //match := re.FindString(text)
    match := re.FindAllString(text, -1)
    fmt.Println(match)
    match_sub := re.FindAllStringSubmatch(text, -1)
    fmt.Println(match_sub)
    for _, m  := range match_sub {
        fmt.Println(m)
    }
}
