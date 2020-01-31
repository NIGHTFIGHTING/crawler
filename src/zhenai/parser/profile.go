package parser

import (
    "engine"
    "strconv"
    "model"
    "regexp"
    //"fmt"
)

// </div><div class="m-btn purple" data-v-8b1eac0c>32岁</div>
// const ageRe = `<div class="m-btn purple" data-v-8b1eac0c>([\d]+)岁</div>`
// 预先编译
var ageRe = regexp.MustCompile(
    //`<div class="m-btn purple"[^>]*>([\d]+)岁</div>`)
    `<div class="m-btn purple"[^>]*>([\d]+)[^<]*</div>`)
// <div class="m-btn purple" data-v-8b1eac0c>离异</div> 
// const marriageRe = `<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`
var marriageRe = regexp.MustCompile(
    `<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)

var heightRe = regexp.MustCompile(
    `<div class="m-btn purple" data-v-8b1eac0c>([^<]+)cm</div>`)
//var weightRe = regexp.MustCompile(
//    `<div class="m-btn purple" data-v-8b1eac0c>([^<]+)cm</div>`)

func ParseProfile(
    contents []byte, name string) engine.ParseResult {
    profile := model.Profile{}
    profile.Name = name
    age, err := strconv.Atoi(
            extractString(contents, ageRe))
    if err != nil {
        profile.Age = age
    }
    height, err := strconv.Atoi(
            extractString(contents, heightRe))
    if err != nil {
        profile.Height = height
    }
//    weight, err := strconv.Atoi(
//            extractString(contents, weightRe))
//    if err != nil {
//        profile.weight = weight
//    }

    profile.Marriage = extractString(
            contents, marriageRe)

    result := engine.ParseResult {
        Items: []interface{}{profile},
    }
    return result

    // re : = regexp.MustCompile(ageRe)
    // match := re.FindSubmatch(contents)
    // if match != nil {
    //     age, err := strconv.Atoi(string(match[1]))
    //     if err != nil {
    //         // user age is age 
    //     }
    // }
}

func extractString(contents []byte, re *regexp.Regexp) string {
    match := re.FindSubmatch(contents)
    if len(match) >= 2 {
        return string(match[1])
    } else {
        return "" 
    }
}
