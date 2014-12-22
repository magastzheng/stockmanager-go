package shseparser

import(
    "entity/stentity"
    "strings"
    "parser/jsonparser"
    "handler/shsehandler"
    //"fmt"
)

type CompanyParser struct {
    Companies []stentity.Company
}

func (p *CompanyParser) Parse(data string) int {
    p.Companies = make([]stentity.Company, 0)
    start := strings.Index(data, "({")
    end := strings.LastIndex(data, "})")
    str := string(data[start+1: end+1])

    //fmt.Println(str)
    handler := shsehandler.NewCompanyHandler()
    parser := jsonparser.NewJsonParser(handler)
    parser.ParseStr(str)
    
    p.Companies = append(p.Companies, handler.Company)

    return 0
}

func NewCompanyParser() *CompanyParser{
    return &CompanyParser{}
}
