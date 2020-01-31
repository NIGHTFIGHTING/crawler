package parser

import (
    "regexp"
    "engine"
)

// view-source:http://www.zhenai.com/zhenghun/aba
// <a href="http://album.zhenai.com/u/1419353990" target="_blank">任我潇洒</a>
const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a>`

func ParseCity(
    contents []byte) engine.ParseResult {
    re := regexp.MustCompile(cityRe)
    matches := re.FindAllSubmatch(contents, -1)

    result := engine.ParseResult{}
    for _, m := range matches {
        // 需要将name拷贝一份
        name := string(m[2])
        result.Items = append(
            result.Items, "User " + string(m[2]))
        result.Requests = append(
            result.Requests, engine.Request{
                Url: string(m[1]),
                // ParserFunc: engine.NilParser,
                // 函数式编程
                ParserFunc: func(
                    c []byte) engine.ParseResult {
                    return ParseProfile(
                        c, name)
                },
            })
    }
    return result
}
