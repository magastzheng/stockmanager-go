package dbentity

const(
    Int = "int"
    Float = "float"
    Varchar = "varchar"
    Date = "date"
    Decimal1 = "decimal1"
    Decimal2 = "decimal2"
    Decimal3 = "decimal3"
    Decimal4 = "decimal4"
    HDecimal2 = "hdecimal2" // for the 20 numbers 2 decimal
)

type DBColumn struct{
   Name string
   Type string
   Maxsize int
   IsNotNull bool
   IsPK bool
   IsAutoIncr bool
}

type DBTable struct{
    TableName string
    Columns []*DBColumn
    Keys []*DBColumn
}

type DBData struct {
    Columns []string
    Rows [][]string
}

type DBExecData struct {
    Rows [][]interface{}
}

