package config

import (
    "encoding/json"
    "fmt"
    //"os"
    //"bytes"
    //"io"
    "io/ioutil"
)

type Category struct {
    Type string
    Class string
    Exchange string
}

type StockSites struct {
    BaseUrl string
    Categories [] Category
}

type StockListConfig struct {
    Sites StockSites
}

func Parse(filename string) StockListConfig {
    chunks, err := ioutil.ReadFile(filename)
    if err != nil {
        panic(err)
    }

    var config StockListConfig
    json.Unmarshal(chunks, &config)
    
    fmt.Println(config)
    return config
}
