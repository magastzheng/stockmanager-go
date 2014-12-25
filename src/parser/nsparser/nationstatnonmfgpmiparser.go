package nsparser

import(
    ns "entity/nsentity"
    "util"
    //"fmt"
)

type NSNonMfgPmiParser struct {
    NSParserBase
    Data []ns.NSNonMfgPmi
}

func (p *NSNonMfgPmiParser) Parse(mapData map[string] string) {
    p.Data = make([]ns.NSNonMfgPmi, 0)
    tempMap := make(map[string]ns.NSNonMfgPmi)

    for k, v := range mapData {
        id, date := p.ParseKey(k)
        
        ms, ok := tempMap[date]
        if !ok {
            ms.Date = date
            tempMap[date] = ms
        }
        
        d := util.ToFloat32(v)
        switch id {
            case "A090201":
                ms.Pmi = d
            case "A090202":
                ms.NewOrder = d
            case "A090203":
                ms.NewExportOrder = d
            case "A090204":
                ms.InHandOrder = d
            case "A090205":
                ms.Inventory = d
            case "A090206":
                ms.ImInputPrice = d
            case "A090207":
                ms.SubscriptionPrice = d
            case "A090208":
                ms.Employees = d
            case "A090209":
                ms.SupplierDeliveryTime = d
            case "A09020A":
                ms.BizActExpectation = d
        }

        tempMap[date] = ms
    }
    
    for _, v := range tempMap {
        p.Data = append(p.Data, v)
    }
}

func NewNSNonMfgPmiParser() *NSNonMfgPmiParser {
    p := new(NSNonMfgPmiParser)
    
    return p
}
