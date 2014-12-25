package account_test

import(
    "testing"
    "excel/account"
    acc "entity/accountentity"
    "fmt"
)

func Test_AccountColumnParser_Parse(t *testing.T){
    filename := "../../resource/account/financialindexdb.xlsx"
    p := account.NewAccountColumnParser()
    p.Parse(filename)
    fmt.Println(len(p.ColumnMap))
    fmt.Println(len(p.ColumnTableMap))
    OutPut_ColumnMap(p.ColumnMap)
    Output_ColumnTableMap(p.ColumnTableMap)
}

func OutPut_ColumnMap(colMap map[string]*acc.Column){
    for k, col := range colMap {
        s := fmt.Sprintf("Key: %s, col: %s, type: %s, maxsz: %d", k, col.Column, col.Type, col.Maxsize)
        fmt.Println(s)
    }
}

func Output_ColumnTableMap(ctMap map[string][]*acc.Column){
    for k, table := range ctMap{
        fmt.Println("Table/Common: ", k, "len:", len(table))

        for _, col := range table {
            s := fmt.Sprintf("Col: %s, type: %s, maxsz: %d", col.Column, col.Type, col.Maxsize)
            fmt.Println(s)
        }
    }
}
