package download

import (
    "net/http"
    "io/ioutil"
    "util"
    "bytes"
    "fmt"
    "github.com/golang/glog"
)

func HttpGet(url string) string {
    resp, err := http.Get(url)
    defer resp.Body.Close()
    //util.CheckError(err)
    if err != nil {
        //TODO: It will use log here
        fmt.Println(err)
        glog.Error(err)
        return ""
    }
    
    body, err := ioutil.ReadAll(resp.Body)
    util.CheckError(err)

    result := string(body)
    return result
}

func HttpPost(url, data string) string {
    buf := []byte(data)
    resp, err := http.Post(url, "application/json", bytes.NewReader(buf))
    //util.CheckError(err)
    defer resp.Body.Close()
    if err != nil {
        //TODO: It will use log here
        fmt.Println(err)
        glog.Error(err)
        return ""
    }

    body, err := ioutil.ReadAll(resp.Body)
    util.CheckError(err)

    result := string(body)
    return result
}
