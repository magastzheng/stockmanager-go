package parser

import(
    "encoding/json"
    "util"
)

type NSBase struct {
    Id string `json: "id"`
    Name string `json: "name"`
}

type NSRegion struct {
    NSBase
    EName string `json: "ename"`
}

type NSIndex struct {
    NSRegion
    IfData string `json: "ifData"`
    IsParent bool `json: "isParent"`
    PId string `json: "pId"`
}

type NSDataIndex struct {
    NSRegion
    Unit string `json: "unit"`
    Eunit string `json: "eunit"`
    Note string `json: "note"`
    Enote string `json: "enote"`
}

type NSValue struct {
    Region []NSRegion `json: "region"`
    Index []NSDataIndex `json: "index"`
    Time []NSBase `json: "time"`
}

type NSData struct {
    TableData map[string] string `json: "tableData"`
    Value NSValue `json: "value"`
}

type NSParser struct{

}

func (p *NSParser) ParseData(data string) NSData {
    var nsdata NSData
    bytes := []byte(data)
    err := json.Unmarshal(bytes, &nsdata)
    if err != nil {
        util.NewLog().Error("Cannot parse nationstat data: ", data)
        util.NewLog().Error(err)
    }

    return nsdata
}

func (p *NSParser) ParseIndex(data string) []NSIndex {
    var nsdata []NSIndex
    bytes := []byte(data)
    err := json.Unmarshal(bytes, &nsdata)
    if err != nil {
        util.NewLog().Error("Cannot parse nationstat index: ", data)
        util.NewLog().Error(err)
    }

    return nsdata
}

func (p *NSParser) ParsePeriod(data string) []NSBase {
    var nsdata []NSBase
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
