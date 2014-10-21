package config

import (
    "testing"
    "fmt"
)

func Test_Parsejson(t *testing.T){
    filename := "stocklist.json"
    config := Parse(filename)
    fmt.Println(config.Sites.BaseUrl)
}
