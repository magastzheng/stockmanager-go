package stockdb_test

import(
    "testing"
    "stockdb"
    "fmt"
)

func Test_SqlGenerator_Insert(t *testing.T){
    cols := []string{"id", "name", "age"}
    g := stockdb.NewSqlGenerator()
    res := g.GenerateInsert("test", cols)
    fmt.Println(res)
}
