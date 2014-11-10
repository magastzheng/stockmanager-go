package stockdb_test

import(
    "testing"
    "stockdb"
    "entity"
    "fmt"
)

func Test_StockHistDataInsert(t *testing.T){
    db := stockdb.NewStockHistDataDB("chinastock")
    data := entity.StockHistData{
        Date: "2014-10-31",
        Open: 3.15,
        Close: 2.46,
        Highest: 5.45,
        Lowest: 1.95,
        Volume: 125897556,
        Money: 45789524,
    }
    
    res := db.Delete("00258", data.Date)
    res = db.Insert("00258", data)
    fmt.Println(res)

    newdata := db.Query("00258", data.Date)
    fmt.Println(newdata)
}
