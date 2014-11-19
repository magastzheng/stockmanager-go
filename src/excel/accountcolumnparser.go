package excel

import(
    "github.com/tealeg/xlsx"
    "entity/dbentity"
    "util"
    "fmt"
)

type AccountColumnParser struct{
    Columns []*dbentity.DBColumn
    logger *util.StockLog
}

func (p *AccountColumnParser)Parse(filename string) {
    p.logger = util.NewLog()
    file, err := xlsx.OpenFile(filename)
    if err != nil{
        p.logger.Error("Cannot open the excel:", filename, err)
        return
    }
    
    for i, sheet := range file.Sheets{
        fmt.Println(i, sheet.Name)   
    }
}

func NewAccountColumnParser() *AccountColumnParser{
    return &AccountColumnParser{}
}

