package config

import(
    "encoding/json"
    "io/ioutil"
    "util"
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
    chunks, err := ioutil.ReadFile(filename)
    util.CheckError(err)

    err = json.Unmarshal(chunks, &m.config)
    util.CheckError(err)
}

func (m *ServiceConfigManager) GetConfig(id, key string) ServiceAPI {
    var item ServiceItem
    isExist := false
    items := m.config.Services

    for _, v := range items {
        if v.Id == id {
            item = v
            isExist = true
            break
        }
    }
    
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
    m := new(ServiceConfigManager)
    m.Parse("serviceconfig.json")

    return m
}
