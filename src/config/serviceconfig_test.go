package config_test

import(
    "testing"
    "config"
)

func Test_ParseServiceConfig(t *testing.T) {
    m := config.NewServiceConfigManager()
    res := m.GetConfig("sina-price", "current")

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
    m.Parse()

    return m
}
