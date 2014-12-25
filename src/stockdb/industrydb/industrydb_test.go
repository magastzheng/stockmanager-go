package industrydb_test

import (
    "testing"
    "entity/xlsentity"
    "fmt"
    "stockdb/industrydb"
)

func Test_IndustryDb_Insert(t *testing.T) {
    db := industrydb.NewIndustryDB("chinastock", "csrcindustry")
    industry := xlsentity.Industry {
        Code: "ATEST",
        Parent: "",
        Name: "测试",
        Name_en: "Test",
    }
    
    res := db.Delete(industry.Code)
    if res = db.Insert(industry); res == -1 {
        t.Error("Cannot insert")
    }
    fmt.Println(res)
}

func Test_IndustryDb_Query(t *testing.T) {
    db := industrydb.NewIndustryDB("chinastock", "csrcindustry")
    
    if industry := db.Query("ATEST"); len(industry.Code) == 0 {
        t.Error("Cannot query")
    }
   
    //fmt.Println(industry)
}
