package excel

import(
    "strings"
    //"fmt"
	"time"
    "bufio"
    "util"
)

const(
	AccDateFormat = "20060102"
)
type AccountParser struct {
    logger *util.StockLog
}

func (p *AccountParser) Init(){
	p.logger = util.NewLog()
}

func (p *AccountParser) Parse(data string) map[string]map[string]float32 {
    
    srd := strings.NewReader(data)
    brd := bufio.NewReader(srd)
    
	//date list
	dates := make([]string, 0)
	// date - index - data mapping
	dataMap := make(map[string]map[string]float32)
	
    //fmt.Println(len(data))
	row := 0
    for line, err := brd.ReadString('\n'); err == nil; line, err = brd.ReadString('\n') {
        line = strings.TrimSuffix(line, " ")
        columns := strings.Split(line, "\t")
		
		if row == 0 {
			// the date row
			for i, col := range columns {
				if i > 0 {
					date := p.FormatDate(col)
					dates = append(dates, date)
				}else{
					dates = append(dates, col)
				}
			}
		} else {
			var key string
			for i, col := range columns {
				if i == 0 {
					key = col
				} else {
					val := util.ToFloat32(col)
					
					if i < len(dates) {
						date := dates[i]
						
						datakeyval, ok := dataMap[date]
						if ok {
							datakeyval[key] = val
							dataMap[date] = datakeyval
						} else {
							datakeyval = make(map[string]float32)
							datakeyval[key] = val
							dataMap[date] = datakeyval
						}
					} else {
						p.logger.Error("Data wrong in the line: ", row, line)
					}	
				}
			}
		}
		row++
    }
	
	return dataMap
}

func (p *AccountParser) FormatDate(date string) string {
	dt, err := time.Parse(AccDateFormat, date)
	if err != nil{
		p.logger.Error("Cannot parse the date: ", date)
		return date
	}
	
	return util.FormatDate(dt)
}

func NewAccountParser() *AccountParser{
	p := new(AccountParser)
	p.Init()
	
    return p
}
