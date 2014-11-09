package parser

import(
    "stockhandler"
    //"io/ioutil"
    "strings"
    "util"
    "fmt"
)

type StockRtParser struct {
    Data stockhandler.StockHistData
}

func (p *StockRtParser) ParseStr(data string) {
    //bytes := []byte(data)
    //var values []string
    //err := json.Unmarshal(bytes, &values)
    //if err != nil {
    //    util.NewLog().Error("Cannot convert json data", data)
    //}
    
    start := strings.Index(data, "\"")
    end := strings.LastIndex(data, "\"")
    validData := data[start+1:end]
    fmt.Println(validData)
    values := strings.Split(validData, ",")
    fmt.Println(values)

    p.Data = stockhandler.StockHistData{}
    for i, v := range values {
        switch i {
            case 1:
                p.Data.Open = util.ToFloat32(v)
            case 3:
                p.Data.Close = util.ToFloat32(v)
            case 4:
                p.Data.Highest = util.ToFloat32(v)
            case 5:
                p.Data.Lowest = util.ToFloat32(v)
            case 8:
                p.Data.Volume = util.ToInt(v)
            case 9:
                p.Data.Money = util.ToInt(v)
            case 30:
                p.Data.Date = v
        }
    }
}

func NewStockRtParser() *StockRtParser{
    p := new(StockRtParser)
    return p
}
