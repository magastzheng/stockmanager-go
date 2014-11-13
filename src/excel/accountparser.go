package excel

import(
    "strings"
    "fmt"
    "bufio"
    "util"
)

type AccountParser struct {
    Head []string
    Title map[string] int
    TableData [][]string

    logger *util.StockLog
}

func (p *AccountParser) Parse(data string) {
    p.logger = util.NewLog()
    srd := strings.NewReader(data)
    brd := bufio.NewReader(srd)
    
    fmt.Println(len(data))
    for line, err := brd.ReadString('\n'); err == nil; line, err = brd.ReadString('\n') {
        fmt.Println("Old:", line)
        line = strings.TrimSuffix(line, " ")
        fmt.Println("New:", line+"|")
        columns := strings.Split(line, "\t")
        //fmt.Println(columns)
        //p.logger.Info(columns)
        str := ""
        for i, col := range columns {
            str = fmt.Sprintf("%d|%v", i, col)
        }
        fmt.Println(str)
        p.TableData = append(p.TableData, columns)
    }
    
    //for i, row := range p.TableData{
        //fmt.Println(i, row)
        //str := fmt.Sprintf("%d\t", i)
        //j := 0
        //for j, col := range row {
            //str += fmt.Sprintf("%d|%v ", j, col)
            //j++
        //}

        //fmt.Println(str)
   // }
}

func NewAccountParser() *AccountParser{
    return &AccountParser{}
}
