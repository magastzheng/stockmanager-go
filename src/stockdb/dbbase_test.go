package stockdb_test

import (
    "testing"
    "stockdb"
    "fmt"
)

func Test_DBBaseQuery(t *testing.T){
    db := new(stockdb.DBBase)
    db.Init("stocktest")
    //query := "select * from stunittest"
    query := "select * from stunittest where id=?"
    //args := make([]interface{}, 0)
    //args = append(args, "000031")
    data := db.Query(query, "000031")
    //data := db.Query(query)
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

func Test_DBBaseExecOnce(t *testing.T){
    db := new(stockdb.DBBase)
    db.Init("stocktest")
    query := "delete from stunittest where id=?"
    db.ExecOnce(query, 3)

    query = "insert stunittest set id=?, name=?, age=?"
    rdata := make([]interface{}, 0)
    rdata = append(rdata, 3)
    rdata = append(rdata, "magast")
    rdata = append(rdata, 34)
    
    db.ExecOnce(query, rdata ...)
}

func Test_DBBaseExec(t *testing.T){
    db := new(stockdb.DBBase)
    db.Init("stocktest")
    query := "insert stunittest set id=?, name=?, age=?"
    data := make([][]interface{}, 0)
    rdata := make([]interface{}, 0)
    rdata = append(rdata, 1)
    rdata = append(rdata, "magast")
    rdata = append(rdata, 31)
    data = append(data, rdata)

    rdata1 := make([]interface{}, 0)
    rdata1 = append(rdata1, 2)
    rdata1 = append(rdata1, "magast1")
    rdata1 = append(rdata1, 30)
    data = append(data, rdata1)
    
    dbdata := stockdb.DBExecData{
        Rows: data,
    }
    db.Exec(query, dbdata)
}

