package download

import (
    "net/http"
    "net/url"
    "io/ioutil"
    "util"
    "bytes"
    //"fmt"
)

func HttpGet(url string) string {
    //fmt.Println(url)
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

func HttpPostForm(uri, query string) string {
    values, err := url.ParseQuery(query)
    if err != nil {
        util.NewLog().Error("URL: ", uri, " cannot parse the form query: ", query, err)
        return ""
    }

    resp, err := http.PostForm(uri, values)
    if resp == nil {
        return ""
    } 
    defer resp.Body.Close()
    if err != nil {
        util.NewLog().Error("URL: ", uri, "Form Data: ", query, err)
        return ""
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        util.NewLog().Error("URL: ", uri, err)
        return ""
    }

    result := string(body)
    return result
}
