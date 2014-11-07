package download

import(
    "net/http"
    "net/url"
    "time"
    "fmt"
    "io/ioutil"
)

func GetRoot() string {
    now := time.Now()
    fmt.Println(now)
    resp, err := http.PostForm("http://data.stats.gov.cn/quotas/getchildren", url.Values{"code":{"A0B"}, "dbcode":{"hgyd"}, "dimension": {"zb"}, "level":{"1"}})
    defer resp.Body.Close()
    if err != nil {
        fmt.Println(err)
        return ""
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return ""
    }

    result := string(body)
    fmt.Println(result)
    return result
}
