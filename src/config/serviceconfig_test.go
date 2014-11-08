package config_test

import(
    "testing"
    "config"
    "fmt"
)

func Test_ParseServiceConfig(t *testing.T) {
    m := config.NewServiceConfigManager()
    res := m.GetConfig("sina-price", "current")
    fmt.Println(res)
}
