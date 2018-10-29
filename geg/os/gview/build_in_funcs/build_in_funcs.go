package main

import (
    "fmt"
    "gitee.com/johng/gf/g"
)

func main() {
    tplContent := `
{{"<div>测试</div>"|text}}
{{"<div>测试</div>"|html}}
{{"&lt;div&gt;测试&lt;/div&gt;"|htmldecode}}
{{"https://gfer.me"|url}}
{{"https://gfer.me"|urldecode}}
{{1540822968 | date "Y-m-d"}}
{{"1540822968" | date "Y-m-d H:i:s"}}
{{compare "A" "B"}}
{{compare "1" "2"}}
{{compare 2 1}}
{{compare 1 1}}
{{"我是中国人" | substr 2 -1}}
{{"我是中国人" | substr 2  2}}
`
    content, err := g.View().ParseContent(tplContent, nil)
    fmt.Println(err)
    fmt.Println(string(content))
}