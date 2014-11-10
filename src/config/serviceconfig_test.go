package config_test

import(
    "testing"
    "config"
    "fmt"
)

func Test_ParseServiceConfig_Service(t *testing.T) {
    m := config.NewServiceConfigManager("serviceconfig.json")
    res, _ := m.GetService("sina-price")
    fmt.Println(res)
    res, _ = m.GetService("nationstat")
    fmt.Println(res)
}

func Test_ParseServiceConfig_Api(t *testing.T) {
    m := config.NewServiceConfigManager("serviceconfig.json")
    res := m.GetApi("sina-price", "current")
    fmt.Println(res)
    res = m.GetApi("sina-price", "realtime")
    fmt.Println(res)
}
