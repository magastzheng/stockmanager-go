package parser

import(
    "entity"
    "util"
    //"strings"
    //"fmt"
)

type NSMSParser struct {
    NSParserBase
    Data []entity.MoneySupply
}

func (p *NSMSParser) Parse(mapData map[string] string) {
    p.Data = make([]entity.MoneySupply, 0)

    tempMap := make(map[string]entity.MoneySupply)
    for k, v := range mapData {
        id, date := p.ParseKey(k)
        
        ms, ok := tempMap[date]
        if !ok {
            ms.Date = date
            tempMap[date] = ms
        }
        
        switch id {
            case "A0B0105":
                ms.M0 = util.ToFloat64(v)
            case "A0B0103":
                ms.M1 = util.ToFloat64(v)
            case "A0B0101":
                ms.M2 = util.ToFloat64(v)
            case "A0B0106":
                ms.M0pct = util.ToFloat32(v)
            case "A0B0104":
                ms.M1pct = util.ToFloat32(v)
            case "A0B0102":
                ms.M2pct = util.ToFloat32(v)
        }

        tempMap[date] = ms
    }
    
    for _, v := range tempMap {
        p.Data = append(p.Data, v)
    }
}

func NewNSMSParser() *NSMSParser {
    p := new(NSMSParser)
    
    return p
}
