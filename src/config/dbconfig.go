package config

import (
    "encoding/json"
    "fmt"
    //"os"
    //"os/exec"
    "runtime"
    "io/ioutil"
    "util"
    "path/filepath"
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
    //fmt.Println(runtime.GOROOT())
    //abs, _ := filepath.Abs(".")
    //fmt.Println("Abs:", abs) 
    //pc, filename, line, ok := runtime.Caller(0)
    //fmt.Println(pc, line, ok, "Base:", filepath.Base(filename))
    //fmt.Println("Dir:", filepath.Dir(filename))
    //fmt.Println(runtime.Caller(0))
    //fmt.Println(os.Getwd())
    //fmt.Println(exec.LookPath(os.Args[0]))
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

func NewDBConfig() *DBConfigManager {
    pc, filename, line, ok := runtime.Caller(0)
    if pc < 0 || line < 0 || !ok {
        fmt.Println("Cannot read the dbconfig.json")
        util.NewLog().Error("Cannot read the file dbconfig.json")
    }

    filename = filepath.Dir(filename) + "/" + "dbconfig.json"
    manager := new(DBConfigManager)
    manager.Parse(filename)

    return manager
}
