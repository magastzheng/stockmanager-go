package industrydb_test

import (
    "testing"
    "entity/xlsentity"
    "fmt"
    "stockdb/industrydb"
)

func Test_StockCategoryDB_Insert(t *testing.T) {
    db := industrydb.NewStockCategoryDB("chinastock", "csrcstockcategory")
    sc := xlsentity.StockCategory {
        Id: "000111",
        Code: "AAA",
    }
    
    res := db.Delete(sc.Id)
    if res = db.Insert(sc); res == -1 {
        t.Error("Cannot insert")
    }
    fmt.Println(res)
}

func Test_StockCategoryDB_TranInsert(t *testing.T) {
    db := industrydb.NewStockCategoryDB("chinastock", "csrcstockcategory")
    
    sc1 := xlsentity.StockCategory {
        Id: "000111",
        Code: "AAA",
    }

    sc2 := xlsentity.StockCategory {
        Id: "000112",
        Code: "AAA",
    }
    
    scs := make([]xlsentity.StockCategory, 0)
    scs = append(scs, sc1, sc2)

    res := db.Delete(sc1.Id)
    res = db.Delete(sc2.Id)
    if res = db.TranInsert(scs); res == -1 {
        t.Error("Cannot insert")
    }
    fmt.Println(res)
}

func Test_StockCategoryDB_Update(t *testing.T) {
    db := industrydb.NewStockCategoryDB("chinastock", "csrcstockcategory")
    
     sc := xlsentity.StockCategory {
        Id: "000111",
        Code: "AAA-Up",
    }
    
    if res := db.Update(sc); res == -1 {
        t.Error("Cannot insert")
    }
}

func Test_StockCategoryDB_Query(t *testing.T) {
    db := industrydb.NewStockCategoryDB("chinastock", "csrcstockcategory")
    
    if sc := db.Query("000111"); len(sc.Code) == 0 {
        t.Error("Cannot query")
    }
   
    //fmt.Println(industry)
}

func Test_StockCategoryDB_QueryAll(t *testing.T) {
    db := industrydb.NewStockCategoryDB("chinastock", "csrcstockcategory")
    
    if scs := db.QueryAll(); len(scs) == 0 {
        t.Error("Cannot query")
    }
   
    //fmt.Println(industry)
}

func Test_StockCategoryDB_QueryStockIds(t *testing.T) {
    db := industrydb.NewStockCategoryDB("chinastock", "csrcstockcategory")
    
    if ids := db.QueryStockIds("AAA"); len(ids) == 0 {
        t.Error("Cannot query")
    }
   
    //fmt.Println(industry)
}

