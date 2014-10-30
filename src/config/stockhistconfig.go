package config

import(
    "encoding/json"
    "io/ioutil"
    "util"
)

type StockHistSourceItem struct {
    Id string `json: "id"`
    Url string `json: "url"`
}

type StockHistSourceConfig struct {
    Sources [] StockHistSourceItem `json: "sources"`
}

type StockHistManager struct {
    Config StockHistSourceConfig
}

func (c *StockHistManager) Parse(filename string) {
    chunks, err := ioutil.ReadFile(filename)
    util.CheckError(err)

    err = json.Unmarshal(chunks, &c.Config)
    util.CheckError(err)
}

func (c *StockHistManager) GetConfig(name string) StockHistSourceItem {
    var item StockHistSourceItem
    items := c.Config.Sources
    for _, v := range items {
        if v.Id == name {
            item = v
            break
        }
    }

    return item
}

func NewStockHistConfig(filename string) *StockHistManager {
    manager := new(StockHistManager)
    manager.Parse(filename)

    return manager
}
