package stsummary_test

import (
    "testing"
    db "stockdb/stsummary"
    "entity/stentity"
    "fmt"
)

func Test_CompanyDBInsert(t *testing.T) {
    //stdb := db.NewCompanyDB("mysql", "root@/chinastock")
    stdb := db.NewCompanyDB("chinastock", "stocksummary")
    c := stentity.Company{}
    c.Code = "1234"
    c.AbbrName = "test"
    c.Name = "test ss"
    c.Name = "test teset"
    c.RegAddr = "dddf fdf dfd"
    c.InceptDate = "2012-12-12"
    c.State = "Guangxi"
    c.City = "Nanning"
    c.Website = "http://www.test.com"

    res := stdb.Delete(c)

    res = stdb.Insert(c)
    fmt.Println(res)

    //res = stdb.Delete(stock)
}

func Test_CompanyDBQuery(t *testing.T) {
    //stdb := stockdb.NewCompanyDB("mysql", "root@/chinastock")
    stdb := db.NewCompanyDB("chinastock", "stocksummary")
    c := stdb.Query("1234")
    fmt.Println(c)
}

