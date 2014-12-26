package download

import (
    "net/http"
    "net/url"
    "io/ioutil"
    "util"
    "bytes"
    //"fmt"
)

func HttpGet(uri string) string {
    //fmt.Println(uri)
    resp, err := http.Get(uri)
    if resp == nil {
        return ""
    }
    defer resp.Body.Close()
    //util.CheckError(err)
    if err != nil {
        util.NewLog().Error("URL: ", uri, err)
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

func HttpGetWithHeader(uri string, header map[string]string) string {
    req, err := http.NewRequest("GET", uri, nil)
	
	for k, v := range header {
		req.Header.Set(k, v)
	}
	if err != nil {
		util.NewLog().Error("URL: ", uri, " header: ", header, err)
	}
	
	client := new(http.Client)
	resp, err := client.Do(req)
	
	if err != nil {
		util.NewLog().Error("URL: ", uri, " header: ", header, err)
	}
	
    if resp == nil {
        return ""
    }
	
    defer resp.Body.Close()
    if err != nil {
        util.NewLog().Error("URL: ", uri, err)
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

func HttpPost(uri, data string) string {
    buf := []byte(data)
    resp, err := http.Post(uri, "application/json", bytes.NewReader(buf))
    if resp == nil {
        return ""
    } 
    defer resp.Body.Close()
    if err != nil {
        util.NewLog().Error("URL: ", uri, "Data: ", data, err)
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
