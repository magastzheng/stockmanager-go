package parser

import(
    "entity"
    "util"
    "strings"
    //"fmt"
)

type NSMSParser struct {
    Data []entity.MoneySupply
}

func (p *NSMSParser) Parse(mapData map[string] string) {
    p.Data = make([]entity.MoneySupply, 0)

    tempMap := make(map[string]entity.MoneySupply)
    for k, v := range mapData {
        id, date := p.ParseKey(k)
        
        ms, ok := tempMap[date]
        if !ok {
            ms = entity.MoneySupply{
                Date: date,
            }

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

//key format as: id_000000_yyyyMM
func (p *NSMSParser) ParseKey(key string) (id, date string) {
    keys := strings.Split(key, "_")
    if len(keys) != 3 {
        util.NewLog().Error("Fail to parse the nation stat key: ", key)
        return 
    }
    
    id = strings.TrimSpace(keys[0])
    t := util.ParseDate(keys[2])

    date = t.Format("2006-01-02")
    //fmt.Println("Key: ", key, " label: ", keys[2], " date: ", date)
    return id, date
}

func NewNSMSParser() *NSMSParser {
    p := new(NSMSParser)
    
    return p
}
