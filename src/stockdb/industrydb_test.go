package stockdb_test

import (
    "testing"
    "excel"
    "fmt"
    "stockdb"
)

func Test_BigIndustryInsert(t *testing.T) {
    db := stockdb.NewIndustryDB("mysql", "root@/chinastock")
    industry := excel.Industry {
        BigCode: "A",
        Name: "测试",
        Name_en: "Test",
    }
    
    res := db.DeleteIndustry(industry.BigCode)
    res = db.InsertIndustry(industry)
    fmt.Println(res)
}
