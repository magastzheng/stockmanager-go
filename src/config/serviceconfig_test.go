package config_test

import(
    "testing"
    "config"
    "fmt"
)

func Test_ParseServiceConfig(t *testing.T) {
    m := config.NewServiceConfigManager("serviceconfig.json")
    res := m.GetConfig("sina-price", "current")
    fmt.Println(res)
    res = m.GetConfig("sina-price", "realtime")
    fmt.Println(res)
}
