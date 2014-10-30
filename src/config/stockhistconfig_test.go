package config_test

import (
    "testing"
    "config"
    //"fmt"
)

func Test_ParseStockHistConfig(t *testing.T){
    filename := "stockhistconfig.json"
    config := config.NewStockHistConfig(filename)

    item := config.GetConfig("sina")

    if item.Id == "sina" {
        t.Log(item.Url)
        t.Log("success")
    }
}
