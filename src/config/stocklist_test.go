package config_test

import (
    "testing"
    "config"
    "fmt"
)

func Test_ParseStocklistjson(t *testing.T){
    filename := "stocklist.json"
    c := config.Parse(filename)
    fmt.Println(c.Sites.BaseUrl)
}
