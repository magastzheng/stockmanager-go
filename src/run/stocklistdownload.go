package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
    "config"
    "parser"
    "stockhandler"
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
        fmt.Print("Network error!\n")
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

func maind() {
    const baseUrl string = "http://app.finance.ifeng.com/hq/list.php"
    const typ = "stock_a"
    const class = "ha"
    s := new(StockDownloader)
    res := s.GetPage(baseUrl, typ, class)

    fmt.Println("length: %d ", len(res))

    fileName := typ + "-" + class + ".dat"
    file, err := os.Create(fileName)
    if err != nil {
        fmt.Println(err.Error())
    }

    defer file.Close()
    file.WriteString(res)
}
