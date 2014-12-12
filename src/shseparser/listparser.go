package shseparser

import(
    "strings"
    "bufio"
    "util"
    "entity"
    "regexp"
    "fmt"
)

type ListParser struct{
    Stocks []entity.Stock
}

func (p* ListParser) Parse(data string) int {
    srd := strings.NewReader(data)
    brd := bufio.NewReader(srd)

    p.Stocks = make([]entity.Stock, 0)
    
    row := 0
    for line, err := brd.ReadString('\n'); err ==  nil; line, err = brd.ReadString('\n') {
        line = strings.TrimSpace(line)

        pattern := "({.*})"
        re, err := regexp.Compile(pattern)
        if err != nil{
            fmt.Println("Cannot compile the regex:", pattern)
            util.NewLog().Error("Cannot compile the regex:", pattern, line)
        }

        stbuf := re.Find([]byte(line))
        stline := string(stbuf)
        if len(stline) > 0 {
            
            arrval := make([]string, 0)
            for pos := strings.Index(stline, "\""); pos != -1; pos = strings.Index(stline, "\"") {
                stline = string(stline[pos + 1:])
                end := strings.Index(stline, "\"")
                val := string(stline[:end])
                stline = string(stline[end+1:])
                arrval = append(arrval, val)
            }
            
            stock := entity.Stock{
                Id: arrval[0],
                Name: arrval[1],
            }

            p.Stocks = append(p.Stocks, stock)
        }

        row++
    }

    return row
}

func NewListParser() *ListParser{
    p := new(ListParser)
    return p
}
