package nsparser

import(
    "encoding/json"
    "util"
    ns "entity/nsentity"
    //"fmt"
)

type NSParser struct{

}

func (p *NSParser) ParseData(data string) ns.NSData {
    var nsdata ns.NSData
    bytes := []byte(data)
    err := json.Unmarshal(bytes, &nsdata)
    if err != nil {
        util.NewLog().Error("Cannot parse nationstat data: ", data)
        util.NewLog().Error(err)
    }

    return nsdata
}

func (p *NSParser) ParseIndex(data string) []ns.NSIndex {
    var nsdata []ns.NSIndex
    bytes := []byte(data)
    err := json.Unmarshal(bytes, &nsdata)
    if err != nil {
        util.NewLog().Error("Cannot parse nationstat index: ", data)
        util.NewLog().Error(err)
    }

    return nsdata
}

func (p *NSParser) ParsePeriod(data string) []ns.NSBase {
    var nsdata []ns.NSBase
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
