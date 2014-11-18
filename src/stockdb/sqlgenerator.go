package stockdb

import(
    "fmt"
)

type SqlGenerator struct{

}

func (g *SqlGenerator)GenerateInsert(table string, columns []string) string {
    sql := "insert " + table + " set"
    
    for i, count := 0, len(columns); i < count; i++ {
        c := columns[i]
        if i < count - 1{
            sql += fmt.Sprintf(" %s=?,", c)
        } else {
            sql += fmt.Sprintf(" %s=?", c)
        }
    }

    return sql
}

func NewSqlGenerator() *SqlGenerator{
    return &SqlGenerator{}
}
