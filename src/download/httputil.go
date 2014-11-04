package download

import (
    "net/http"
    "io/ioutil"
    "util"
    "bytes"
)

func HttpGet(url string) string {
    resp, err := http.Get(url)
    util.CheckError(err)
    defer resp.Body.Close()
    
    body, err := ioutil.ReadAll(resp.Body)
    util.CheckError(err)

    result := string(body)
    return result
}

func HttpPost(url, data string) string {
    buf := []byte(data)
    resp, err := http.Post(url, "application/json", bytes.NewReader(buf))
    util.CheckError(err)
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    util.CheckError(err)

    result := string(body)
    return result
}
