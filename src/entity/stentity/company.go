package stentity

import(
    "fmt"
)

type Company struct {
    Code string
    AbbrName string
    Name string
    Name_en string
    RegAddr string
    InceptDate string
    TotalShares_A int
    FlowShares_A int
    Code_B string
    InceptDate_B string
    TotalShares_B int
    FlowShares_B int
    Region string
    State string
    City string
    Industry string
    Website string
}

func (c *Company) ToString() string {
    format := "Code: %s, Abbr: %s, Name: %s, Name_en: %s, RegAddr: %s, Incept: %s, TotalA: %d, FlowA: %d, CodeB: %s, InceptB: %s, TotalB: %d, FlowB: %d, Region: %s, State: %s, City: %s, Industry: %s, Web: %s"
    s := fmt.Sprintf(format, c.Code, c.AbbrName, c.Name, c.Name_en, c.RegAddr, c.InceptDate, c.TotalShares_A, c.FlowShares_A, c.Code_B, c.InceptDate_B, c.TotalShares_B, c.FlowShares_B, c.Region, c.State, c.City, c.Industry, c.Website)
    return s
}
