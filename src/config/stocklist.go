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
    //file, err := os.Open(filename)
    //if err != nil {
    //    panic(err)
    //}
    //defer file.Close()

    //chunks := bytes.NewBuffer(nil)
    //io.Copy(chunks, file)
    chunks, err := ioutil.ReadFile(filename)
    if err != nil {
        panic(err)
    }

    var config StockListConfig
    json.Unmarshal(chunks, &config)
    
    fmt.Println(config)
    return config
}
