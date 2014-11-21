package config

import (
    "encoding/json"
    "fmt"
    "os"
    "os/exec"
    "io/ioutil"
    "util"
)

type DBItem struct {
    Name string `json: "name"`
    Dbtype string `json:"dbtype"`
    Dbcon string `json:"dbcon"`
}

type DBConfig struct {
    DBItems [] DBItem `json:"dbitems"`
}

type DBConfigManager struct {
    Config DBConfig
}

func (c *DBConfigManager) Parse(filename string) {
    
    fmt.Println(exec.LookPath(os.Args[0]))
    chunks, err := ioutil.ReadFile(filename)
    util.CheckError(err)
   
    err = json.Unmarshal(chunks, &c.Config)
    util.CheckError(err) 
}

func (c *DBConfigManager) GetConfig(name string) DBItem {
    var dbitem DBItem
    items := c.Config.DBItems
    for _, v := range items {
        if v.Name == name {
            dbitem = v
            break
        }
    }

    return dbitem
}

func NewDBConfig(filename string) *DBConfigManager {
    manager := new(DBConfigManager)
    manager.Parse(filename)

    return manager
}
