package config

import(
    "encoding/json"
    "io/ioutil"
    "util"
    "runtime"
    "path/filepath"
)

type Exchange struct{
    Code string     `json: "code"`
    Region string   `json: "region"`
    Name string     `json: "name"`
}

type CountryExchange struct{
    Id string               `json: "id"`
    Exchanges []Exchange    `json: "exchanges"`
}

type ExchangeConfig struct{
    Countries []CountryExchange `json: "countries"`
}

type ExchangeConfigManager struct{
    Config ExchangeConfig
}

func (c *ExchangeConfigManager) Parse(filename string){
    chunks, err := ioutil.ReadFile(filename)
    if err != nil {
        util.NewLog().Error("Cannot parse the file:", filename, err)
    }

    err = json.Unmarshal(chunks, &c.Config)
    if err != nil {
        util.NewLog().Error("Cannot convert the json into entity", err)
    }
}

func (c *ExchangeConfigManager) GetExchange(country, region string) (exch Exchange, ok bool){
    exchcountry, isExist := c.GetCountry(country)
    if isExist {
        for _, v := range exchcountry.Exchanges {
            if v.Region == region{
                exch = v
                ok = true
                break
            }
        }
    }

    return exch, ok
}

func (c *ExchangeConfigManager) GetCountry(country string) (countryexch CountryExchange, ok bool){
    countries := c.Config.Countries
    for _, v := range countries{
        if v.Id == country {
            countryexch = v
            ok = true
            break
        }
    }

    return countryexch, ok
}

func NewExchangeConfigManager() *ExchangeConfigManager{
    pc, filename, line, ok := runtime.Caller(0)
    if pc < 0 || line < 0 || !ok {
        util.NewLog().Error("Cannot read the file exchangeconfig.json")
    }
    filename = filepath.Dir(filename) + "/" + "exchangeconfig.json"

    m := new(ExchangeConfigManager)
    m.Parse(filename)

    return m
}

