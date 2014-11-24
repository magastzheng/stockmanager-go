package config

import(
    "encoding/json"
    "io/ioutil"
    "util"
    "runtime"
    "path/filepath"
    "fmt"
    //"os"
)

type ServiceAPI struct {
    Key string `json: "key"`
    Method string `json: "method"`
    Uri string  `json: "uri"`
    Data string `json: "data"`
}

type ServiceItem struct {
    Id string   `json: "id"`
    Host string `json: "host"`
    Apis [] ServiceAPI  `json: "apis"`
}

type ServiceConfig struct {
    Services [] ServiceItem `json: "services"`
}

type ServiceConfigManager struct {
    config ServiceConfig
}

func (m *ServiceConfigManager) Parse(filename string) {
    //dir, err := os.Getwd()
    //if err != nil {
    //    fmt.Println(err)
    //}
    //fmt.Println("dir:", dir)

    chunks, err := ioutil.ReadFile(filename)
    //util.CheckError(err)
    if err != nil {
        //fmt.Println("Cannot read file:", filename, err)
        panic(err)
    }

    err = json.Unmarshal(chunks, &m.config)
    util.CheckError(err)
}

func (m *ServiceConfigManager) GetService(id string) (ServiceItem, bool) {
    var item ServiceItem
    var ok bool = false
    items := m.config.Services
    
    for _, v := range items {
        if v.Id == id {
            item = v
            ok = true
            break
        }
    }

    return item, ok
}

func (m *ServiceConfigManager) GetApi(id, key string) ServiceAPI {
    item, isExist := m.GetService(id)
    
    var api ServiceAPI
    if isExist {
        for _, v := range item.Apis {
            if v.Key == key {
                api = v
                api.Uri = item.Host + api.Uri
            }
        }
    }

    return api
}

func NewServiceConfigManager() *ServiceConfigManager{
    pc, filename, line, ok := runtime.Caller(0)
    if pc < 0 || line < 0 || !ok {
        fmt.Println("Cannot read the serviceconfig.json")
        util.NewLog().Error("Cannot read the file serverconfig.json")
    }
    filename = filepath.Dir(filename) + "/" + "serviceconfig.json"

    m := new(ServiceConfigManager)
    //m.Parse("./serviceconfig.json")
    m.Parse(filename)

    return m
}
