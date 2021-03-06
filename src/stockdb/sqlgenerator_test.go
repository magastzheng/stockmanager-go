package stockdb_test

import(
    "testing"
    "stockdb"
    "entity/dbentity"
    "fmt"
)

func Test_SqlGenerator_Insert(t *testing.T){
    cols := []string{"id", "name", "age"}
    g := stockdb.NewSqlGenerator()
    res := g.GenerateInsert("test", cols)
    fmt.Println(res)
}

func Test_SqlGenerator_Create(t *testing.T){
    col1 := dbentity.DBColumn{
        Name: "id",
        Type: "int",
        IsPK: true,
        IsNotNull: true,
    }

    col2 := dbentity.DBColumn{
        Name: "name",
        Type: "varchar",
        Maxsize: 20,
        IsNotNull: true,
    }

    col3 := dbentity.DBColumn{
        Name: "age",
        Type: "float",
        Maxsize: 10,
    }


    cols := make([]*dbentity.DBColumn, 0)
    cols = append(cols, &col1, &col2, &col3)
    
    table := dbentity.DBTable{
        TableName: "test_generate_auto",
        Columns: cols,
    }

    g := stockdb.NewSqlGenerator()
    res := g.GenerateCreate(table)

    fmt.Println(res)
}
