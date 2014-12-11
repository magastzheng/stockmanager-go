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
    
}

func (p* ListParser) Parse(data string){
    srd := strings.NewReader(data)
    brd := bufio.NewReader(srd)

    stlist := make([]entity.StockItem, 0)
    
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
            //codepattern := `val:\"\d+\"`
            //namepattern := `val2:\"\.+\"`
            //recode, _ := regexp.Compile(codepattern)
            //codebuf := recode.Find([]byte(line))

            //rename, _ := regexp.Compile(namepattern)
            //namebuf := rename.Find([]byte(line))

            //fmt.Println(string(codebuf), string(namebuf))
            
            for pos := strings.Index(stline, "\""); pos > 0; pos = strings.Index(stline, "\"") {
                stline = string(stline[pos + 1:])
                end := strings.Index(stline, "\"")
                val := string(stline[:end])
                stline = string(stline[end+1:])
                fmt.Println(val, pos, end, stline)
            }
        }

        row++
    }

    fmt.Println(stlist)
}

func NewListParser() *ListParser{
    p := new(ListParser)
    return p
}
