package download

import (
    "net/http"
    "io/ioutil"
)

//http://app.finance.ifeng.com/hq/list.php
//http://app.finance.ifeng.com/hq/list.php?type=stock_a&class=ha
//class can be: ha, sa, gem

type StockDownloader struct{
	id string
}

func (s *StockDownloader) GetUrl(baseUrl, typ, class string) string {
    return  baseUrl+"?type="+typ+"&class="+class
}

func (s *StockDownloader) GetPage(baseUrl, typ, class string) string {
    url := s.GetUrl(baseUrl, typ, class)
    resp, err := http.Get(url)
    if err != nil {
        //fmt.Print("Network error!\n")
        panic(err)
    }
    
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    result := string(body)
    //fmt.Println(result)
    return result
}

func NewDownloader() *StockDownloader {
    return &StockDownloader{}
}
