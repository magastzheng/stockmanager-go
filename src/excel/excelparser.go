package excel

import(
    "github.com/tealeg/xlsx"
    "strconv"
    //"fmt"
)

type MinorIndustry struct {
    MinorCode int
    BigCode string
    Name string
    Name_en string
}

type Industry struct {
    BigCode string
    Name string
    Name_en string
}

type RowInfo struct {
    StockId string
    StockName string
    StockName_en string
    Exchange string
    BigCode string
    BigName string
    BigName_en string
    MinorCode int
    MinorName string
    MinorName_en string
}

type IndustryParser struct {
    MinorMap map[int] MinorIndustry
    BigMap map[string] Industry
    Rows []*RowInfo
}

func (p *IndustryParser) Parse(filename string) {
    file, err := xlsx.OpenFile(filename)
    if err != nil {
        panic(err)
    }

    for _, sheet := range file.Sheets {
        p.Rows = make([]*RowInfo, len(sheet.Rows))
        p.MinorMap = make(map[int] MinorIndustry)
        p.BigMap = make(map[string] Industry)
        for ridx, row := range sheet.Rows {
            if ridx == 0 {
                continue
            }
            
            rowInfo := new(RowInfo)
            for cidx, cell := range row.Cells {
                value := cell.String()
                switch cidx {
                    case 1:
                        rowInfo.StockId = value
                    case 2:
                        rowInfo.StockName = value
                    case 3:
                        rowInfo.StockName_en = value
                    case 4:
                        rowInfo.Exchange = value
                    case 5:
                        rowInfo.BigCode = value
                    case 6:
                        rowInfo.BigName = value
                    case 7:
                        rowInfo.BigName_en = value
                    case 8:
                        minorCode, err := strconv.Atoi(value)
                        if err != nil {
                            panic(err)
                        }

                        rowInfo.MinorCode = minorCode
                    case 9:
                        rowInfo.MinorName = value
                    case 10:
                        rowInfo.MinorName_en = value
                }//end of swtich
            }

            p.Rows = append(p.Rows, rowInfo)
            if _, ok := p.MinorMap[rowInfo.MinorCode]; !ok {
                minorInds := MinorIndustry{
                    MinorCode : rowInfo.MinorCode,
                    BigCode : rowInfo.BigCode,
                    Name : rowInfo.MinorName,
                    Name_en : rowInfo.MinorName_en,
                }

                p.MinorMap[rowInfo.MinorCode] = minorInds
            }

            if _, ok := p.BigMap[rowInfo.BigCode]; !ok {
                bigInds := Industry {
                    BigCode : rowInfo.BigCode,
                    Name : rowInfo.BigName,
                    Name_en : rowInfo.BigName_en,
                }

                p.BigMap[rowInfo.BigCode] = bigInds
            }
        }
    }
}
