package config_test

import (
    "testing"
    "config"
    //"fmt"
)

func Test_ParseDBConfigjson(t *testing.T){
    filename := "dbconfig.json"
    config := config.NewDBConfig(filename)

    item := config.GetConfig("chinastock")

    if item.Dbtype == "mysql" && item.Dbcon == "root@/chinastock" {
        t.Log("success")
    }
}
