package nsentity

type NSBase struct {
    Id string `json: "id,string"`
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
    PId string `json: "pId,string"`
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
