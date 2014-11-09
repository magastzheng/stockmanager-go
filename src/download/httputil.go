package download

import (
    "net/http"
    "io/ioutil"
    "util"
    "bytes"
    //"fmt"
)

func HttpGet(url string) string {
    resp, err := http.Get(url)
    if resp == nil {
        return ""
    }
    defer resp.Body.Close()
    //util.CheckError(err)
    if err != nil {
        util.NewLog().Error("URL: ", url, err)
        return ""
    }
    
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        util.NewLog().Error("URL: ", url, err)
        return ""
    }

    result := string(body)
    return result
}

func HttpPost(url, data string) string {
    buf := []byte(data)
    resp, err := http.Post(url, "application/json", bytes.NewReader(buf))
    if resp == nil {
        return ""
    } 
    defer resp.Body.Close()
    if err != nil {
        util.NewLog().Error("URL: ", url, "Data: ", data, err)
        return ""
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        util.NewLog().Error("URL: ", url, err)
        return ""
    }

    result := string(body)
    return result
}
