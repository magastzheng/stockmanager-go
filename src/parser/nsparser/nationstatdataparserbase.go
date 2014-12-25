package nsparser

import(
    "util"
    "strings"
    //"fmt"
)

type NSParserBase struct {
}

func (p *NSParserBase) Parse(mapData map[string] string) {
    
}

//key format as: id_000000_yyyyMM
func (p *NSParserBase) ParseKey(key string) (id, date string) {
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
