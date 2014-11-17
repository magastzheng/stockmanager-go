package stockdb_test

import (
    "testing"
    "stockdb"
    "fmt"
)

func Test_DBBaseQuery(t *testing.T){
    db := new(stockdb.DBBase)
    db.Init("chinastock")
    query := "select * from stocklist"
    //query := "select * from stocklist where id=?"
    //args := make([]interface{}, 0)
    //args = append(args, "000031")
    //data := db.Query(query, "000031")
    data := db.Query(query)
//    data := db.Query(query, args)
    fmt.Println(data.Columns)
    fmt.Print("\t\t")
    for _, c := range data.Columns {
        fmt.Print(c, "\t")
    }
    fmt.Print("\n")
    for i, row := range data.Rows{
        fmt.Print(i, "\t")
        for _, c := range row {
            fmt.Print(c, "\t")
        }
        fmt.Print("\n")
    }
}

