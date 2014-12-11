package config_test

import(
    "testing"
    "config"
    "fmt"
)

func Test_ExchangeConfig_Parse(t *testing.T){
    c := config.NewExchangeConfigManager()
    exchange, ok := c.GetExchange("CHS", "Shanghai")

    fmt.Println(exchange, ok)
}
