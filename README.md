# baidu_translate_go
Baidu translate golang sdk / 百度翻译 golang SDK

### How to use.
Get your `appid` and `appkey` from baidu developer console. https://fanyi-api.baidu.com/manage/developer

```go
package main

import (
    "fmt"
    "github.com/shenjinti/baidu_translate_go"
)

func main() {
    appId := ""
    appKey := ""

    bt := baidutranslate.NewBaiduTranslate(appId, appKey)
    v, _ := bt.Text("auto", "pt", "Hello world")
    fmt.Println("result", v)
}

```

