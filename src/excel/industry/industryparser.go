package industry

import(
    "github.com/tealeg/xlsx"
	"util"
    "entity/xlsentity"
    "strings"
    //"fmt"
)

type IndustryRow struct {
    Column1 string
    Column2 string
    Column3 string
}

type IndustryParser struct {
    MinorMap map[string] xlsentity.Industry
    BigMap map[string] xlsentity.Industry
    NewRows []*IndustryRow
    OldRows []*IndustryRow
}

func (p *IndustryParser) Parse(filename string) {
    file, err := xlsx.OpenFile(filename)
    if err != nil {
        panic(err)
    }
	
	var bigCode string
    for i, sheet := range file.Sheets {
        rowlen := len(sheet.Rows)
        switch i {
            case 0:
                p.NewRows = make([]*IndustryRow, rowlen)
                p.MinorMap = make(map[string] xlsentity.Industry)
                p.BigMap = make(map[string] xlsentity.Industry)
				bigCode = p.ParseIndustryRow(bigCode, true, sheet.Rows)
            case 1:
				
            case 2:
                p.OldRows = make([]*IndustryRow, rowlen)
				bigCode = p.ParseIndustryRow(bigCode, false, sheet.Rows)
            case 3:
        }
    }
}

func (p *IndustryParser) ParseIndustryRow(bigCode string, newStd bool, rows []*xlsx.Row) string {
    for ridx, row := range rows {
		if ridx == 0 {
			continue
		}
		
		rowInfo := new(IndustryRow)
		for cidx, cell := range row.Cells {
			value := strings.TrimSpace(cell.String())
			switch cidx {
				case 0:
					rowInfo.Column1 = value
				case 1:
					rowInfo.Column2 = value
				case 2:
					rowInfo.Column3 = value
			}//end of swtich
		} //end of cells
	    
        //fmt.Println(rowInfo)
		if newStd {
			p.NewRows = append(p.NewRows, rowInfo)
		} else {
			p.OldRows = append(p.OldRows, rowInfo)
		}
		
		if util.IsStringNotEmpty(rowInfo.Column1) {
			bigCode = rowInfo.Column1
			if _, ok := p.BigMap[bigCode]; !ok {
				bigInds := xlsentity.Industry {
					Code : bigCode,
				}
				
				if newStd {
					bigInds.Name = rowInfo.Column2
				} else {
					bigInds.Name_en = rowInfo.Column2
				}
					
				p.BigMap[bigCode] = bigInds
			} else {
				bigInds := p.BigMap[bigCode]
				if newStd {
					bigInds.Name = rowInfo.Column2
				} else {
					bigInds.Name_en = rowInfo.Column2
				}
                p.BigMap[bigCode] = bigInds
			}
		} else {
            if !util.IsStringNotEmpty(rowInfo.Column2) {
                return bigCode
            }

            minorCode := rowInfo.Column2
			if _, ok := p.MinorMap[minorCode]; !ok {
				minorInds := xlsentity.Industry{
					Code : minorCode,
					Parent : bigCode,
					//Name : rowInfo.MinorName,
				}
				
				if newStd {
					minorInds.Name = rowInfo.Column3
				} else {
					minorInds.Name_en = rowInfo.Column3
				}
				
				p.MinorMap[minorCode] = minorInds
			} else {
				minorInds := p.MinorMap[minorCode]

				if newStd {
					minorInds.Name = rowInfo.Column3
				} else {
					minorInds.Name_en = rowInfo.Column3
				}
                
                p.MinorMap[minorCode] = minorInds
			}
		}
	}
	
	return bigCode
}

func NewIndustryParser(filename string) *IndustryParser {
    parser := new(IndustryParser)
    parser.Parse(filename)

    return parser
}
