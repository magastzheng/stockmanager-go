package parser

import(
    "encoding/json"
    "util"
    "entity"
)

type NSParser struct{

}

func (p *NSParser) ParseData(data string) entity.NSData {
    var nsdata entity.NSData
    bytes := []byte(data)
    err := json.Unmarshal(bytes, &nsdata)
    if err != nil {
        util.NewLog().Error("Cannot parse nationstat data: ", data)
        util.NewLog().Error(err)
    }

    return nsdata
}

func (p *NSParser) ParseIndex(data string) []entity.NSIndex {
    var nsdata []entity.NSIndex
    bytes := []byte(data)
    err := json.Unmarshal(bytes, &nsdata)
    if err != nil {
        util.NewLog().Error("Cannot parse nationstat index: ", data)
        util.NewLog().Error(err)
    }

    return nsdata
}

func (p *NSParser) ParsePeriod(data string) []entity.NSBase {
    var nsdata []entity.NSBase
    bytes := []byte(data)
    err := json.Unmarshal(bytes, &nsdata)
    if err != nil {
        util.NewLog().Error("Cannot parse nationstat period: ", data)
        util.NewLog().Error(err)
    }

    return nsdata
}

func NewNSParser() *NSParser{
    p := new(NSParser)
    return p
}
