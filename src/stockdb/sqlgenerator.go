package stockdb

import(
    "fmt"
    "bytes"
    "entity/dbentity"
)

type SqlGenerator struct{

}

func (g *SqlGenerator) getSqlType(ctype string, maxsize int) string {
    switch ctype{
        case dbentity.Int:
            return "int"
        case dbentity.Float:
            return "float"
        case dbentity.Varchar:
            return fmt.Sprintf("varchar(%d)", maxsize)
        case dbentity.Date:
            return "date"
        case dbentity.Decimal1:
            return fmt.Sprintf("decimal(%d,%d)", maxsize, 1)
        case dbentity.Decimal2:
            return fmt.Sprintf("decimal(%d,%d)", maxsize, 2)
        case dbentity.Decimal3:
            return fmt.Sprintf("decimal(%d,%d)", maxsize, 3)
        case dbentity.Decimal4:
            return fmt.Sprintf("decimal(%d,%d)", maxsize, 4)
        case dbentity.HDecimal2:
            return fmt.Sprintf("decimal(%d,%d)", maxsize, 2)
    }

    return fmt.Sprintf("char(%d)", maxsize)
}

func (g *SqlGenerator) GenerateCreate(table dbentity.DBTable) string {
    s := bytes.Buffer{}
    s.WriteString("create table ")
    s.WriteString(table.TableName)
    s.WriteString(" ( ")
    
    for i, col := range table.Columns {
        if i > 0 {
            s.WriteString(", ")
        }

        stype := g.getSqlType(col.Type, col.Maxsize)
        s.WriteString(fmt.Sprintf("%s %s", col.Name, stype))
        
        if col.IsPK || col.IsNotNull {
            s.WriteString(" not null")
        }

        if col.IsPK && len(table.Keys) == 1{
            s.WriteString(" primary key")
        }

        if col.IsAutoIncr {
            s.WriteString(" auto_increment")
        }
    }

    if len(table.Keys) > 1{
        s.WriteString(", primary key (")
        for i, k := range table.Keys{
            if i > 0{
                s.WriteString(", ")
            }
            s.WriteString(k.Name)
        }

        s.WriteString(")")
    }

    s.WriteString(")")

    return s.String()
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
