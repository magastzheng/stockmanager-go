package download

import (
    "net/http"
    "io/ioutil"
    "util"
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
