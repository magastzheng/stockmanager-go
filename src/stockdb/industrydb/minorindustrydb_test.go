package stockdb_test

import (
    "testing"
    "excel"
    "fmt"
    "stockdb"
)

func Test_BigIndustryInsert(t *testing.T) {
    db := stockdb.NewMinorIndustryDB("chinastock")
    industry := excel.MinorIndustry {
        MinorCode: 12,
        BigCode: "A",
        Name: "测试",
        Name_en: "Test",
    }
    
    res := db.DeleteIndustry(industry.MinorCode)
    res = db.InsertIndustry(industry)
    fmt.Println(res)
}
