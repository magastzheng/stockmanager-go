package shseparser

import{
    "entity/stentity"
    "strings"
}

type CompanyParser struct {
    Companies []stentity.Company
}

fund (p *CompanyParser) Parse(data string) int {
    start := data.Index(data, "({")
    end := data.LastIndex(data, "})")
    
}
