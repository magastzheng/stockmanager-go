package excel

import(
    "github.com/tealeg/xlsx"
    acc "entity/accountentity"
    "util"
    "fmt"
    "strings"
)

type AccountColumnParser struct{
    //use to create the database table, the key is table/common name
    CategoryColumnMap map[string][]*acc.Column
    //use to parse the excel/html data, where the key is "name" column in excel
	//It will be a Chinese name as utf-8 encoding in our definition.
    ColumnMap map[string]*acc.Column
    logger *util.StockLog
}

func (p *AccountColumnParser)Parse(filename string) {
    p.logger = util.NewLog()
    file, err := xlsx.OpenFile(filename)
    if err != nil{
        p.logger.Error("Cannot open the excel:", filename, err)
        return
    }
    p.CategoryColumnMap = make(map[string][]*acc.Column)
    p.ColumnMap = make(map[string]*acc.Column)
    for i, sheet := range file.Sheets{
        fmt.Println(i, sheet.Name)
        p.logger.Info("Start to parse sheet: ", i, sheet.Name)
        //rowlen := len(sheet.Rows)
        
        p.ParseRow(sheet.Rows)
    }
}

func (p *AccountColumnParser) ParseRow(rows []*xlsx.Row) {
    columnMap := make(map[string] int)
    isCommon := false
    isTable := false
    isNormal := false
    
    var parentColName string
    var nmidx, nmeidx, colidx, typeidx, maxszidx int
    for ridx, row := range rows{
        if ridx == 0 {
            //parse the header to init columnMap
            for cidx, cell := range row.Cells {
                value := strings.TrimSpace(cell.String())
                columnMap[value] = cidx
            }

            nmidx = columnMap["name"]
            nmeidx = columnMap["name_en"]
            colidx = columnMap["column"]
            typeidx = columnMap["type"]
            maxszidx = columnMap["maxsize"]
        } else {
            cols := make([]string, 0)
            for _, cell := range row.Cells {
                value := strings.TrimSpace(cell.String())
                cols = append(cols, value)
            }
            
            size := 0
            if len(cols) > maxszidx {
                size = util.ToInt(cols[maxszidx])
            }

            column := acc.Column{
                Name: cols[nmidx],
                Name_en: cols[nmeidx],
                Column: cols[colidx],
                Type: cols[typeidx],
                Maxsize: size,
            }

            if column.Type == acc.Common {
                isCommon = true
                isTable = false
                isNormal = false

                parentColName = column.Column
            } else if column.Type == acc.Table {
                isCommon = false
                isTable = true
                isNormal = false

                parentColName = column.Column
            } else {
                isNormal = true
            }
            
            if (isCommon || isTable) && isNormal {
                ccols, ok := p.CategoryColumnMap[parentColName]
                if !ok {
                    ccols = make([]*acc.Column, 0)
                    ccols = append(ccols, &column)
                    p.CategoryColumnMap[parentColName] = ccols
                } else {
                    ccols = append(ccols, &column)
                    p.CategoryColumnMap[parentColName] = ccols
                }
            }

            if isNormal {
                p.ColumnMap[column.Name] = &column
            }
        }
    }
}




func NewAccountColumnParser() *AccountColumnParser{
    return &AccountColumnParser{}
}

