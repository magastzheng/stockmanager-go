package download_test

import (
    "testing"
    "download"
)

func Test_HttpGet(t *testing.T) {
    result := download.HttpGet("http://www.baidu.com")
    if len(result) > 0 {
        t.Log("Success")
    }
}

func Test_HttpPost(t *testing.T) {
    url := "http://awdqa85.morningstar.com/setting/v1/globalsetting/5514D458-4ED4-4CB0-88E2-3C014B63D4EC/2262E3B5-8F49-49C2-9A39-5099FC3B1435/general/test"
    json := `{"Key": "test","Id": "test","Value": "testv", "Name": "testn", "Disabled": "false"}`

    result := download.HttpPost(url, json)
    if len(result) > 0 {
        t.Log(result)
    } else {
        t.Error("Cannot post")
    }
}

func Test_HttpPostForm(t *testing.T) {
    url := "http://data.stats.gov.cn/quotas/getchildren"
    query := "code=A0B&level=1&dbcode=hgyd&dimension=zb"
    result := download.HttpPostForm(url, query)
    if len(result) > 0 {
        t.Log(result)
    } else {
        t.Error("Cannot post form")
    }
}
