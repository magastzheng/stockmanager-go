package industry

import(
    "github.com/tealeg/xlsx"
    "entity/xlsentity"
    //"strconv"
    //"fmt"
)

type RowInfo struct {
    StockId string
    StockName string
    StockName_en string
    Exchange string
    BigCode string
    BigName string
    BigName_en string
    MinorCode string
    MinorName string
    MinorName_en string
}

type StockIndustryParser struct {
    MinorMap map[string] xlsentity.Industry
    BigMap map[string] xlsentity.Industry
    Rows []*RowInfo
    Scs []xlsentity.StockCategory
}

func (p *StockIndustryParser) Parse(filename string) {
    file, err := xlsx.OpenFile(filename)
    if err != nil {
        panic(err)
    }

    p.MinorMap = make(map[string] xlsentity.Industry)
    p.BigMap = make(map[string] xlsentity.Industry)
    p.Rows = make([]*RowInfo, 0)
    p.Scs = make([]xlsentity.StockCategory, 0)

    for _, sheet := range file.Sheets {
       
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
                        rowInfo.MinorCode = value
                    case 9:
                        rowInfo.MinorName = value
                    case 10:
                        rowInfo.MinorName_en = value
                }//end of swtich
            }

            p.Rows = append(p.Rows, rowInfo)
            
            sc := xlsentity.StockCategory {
                Id: rowInfo.StockId,
                Code: rowInfo.MinorCode,
            }

            p.Scs = append(p.Scs, sc)

            if _, ok := p.MinorMap[rowInfo.MinorCode]; !ok {
                minorInds := xlsentity.Industry{
                    Code : rowInfo.MinorCode,
                    Parent : rowInfo.BigCode,
                    Name : rowInfo.MinorName,
                    Name_en : rowInfo.MinorName_en,
                }

                p.MinorMap[rowInfo.MinorCode] = minorInds
            }

            if _, ok := p.BigMap[rowInfo.BigCode]; !ok {
                bigInds := xlsentity.Industry {
                    Code : rowInfo.BigCode,
                    Name : rowInfo.BigName,
                    Name_en : rowInfo.BigName_en,
                }

                p.BigMap[rowInfo.BigCode] = bigInds
            }
        }
    }
}

func NewStockIndustryParser(filename string) *StockIndustryParser {
    parser := new(StockIndustryParser)
    parser.Parse(filename)

    return parser
}
