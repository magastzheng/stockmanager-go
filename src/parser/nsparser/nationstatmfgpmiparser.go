package nsparser

import(
    ns "entity/nsentity"
    "util"
    //"fmt"
)

type NSMfgPmiParser struct {
    NSParserBase
    Data []ns.NSMfgPmi
}

func (p *NSMfgPmiParser) Parse(mapData map[string] string) {
    p.Data = make([]ns.NSMfgPmi, 0)

    tempMap := make(map[string]ns.NSMfgPmi)
    for k, v := range mapData {
        id, date := p.ParseKey(k)
        
        ms, ok := tempMap[date]
        if !ok {
            ms.Date = date
            tempMap[date] = ms
        }
        
        d := util.ToFloat32(v)
        switch id {
            case "A090101":
                ms.Pmi = d
            case "A090102":
                ms.Production = d
            case "A090103":
                ms.NewOrder = d
            case "A090104":
                ms.NewExportOrder = d
            case "A090105":
                ms.InHandOrder = d
            case "A090106":
                ms.Inventory = d
            case "A090107":
                ms.PurchasingVolume = d
            case "A090108":
                ms.Import = d
            case "A090109":
                ms.MainRawMaterial = d
            case "A09010A":
                ms.RawMaterialInventory = d
            case "A09010B":
                ms.Employees = d
            case "A09010C":
                ms.SupplierDeliveryTime = d
            case "A09010D":
                ms.PbizActExpectation = d
        }
        
        //fmt.Println(ms.Date)
        tempMap[date] = ms
    }

    for _, v := range tempMap {
        p.Data = append(p.Data, v)
    }
}

func NewNSMfgPmiParser() *NSMfgPmiParser {
    p := new(NSMfgPmiParser)
    
    return p
}
