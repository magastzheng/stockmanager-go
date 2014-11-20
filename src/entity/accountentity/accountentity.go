package accountentity

const(
    Common = "common"
    Table = "table"
)

type Column struct{
    Name string
    Name_en string
    Column string //database column name
    Type string
    Maxsize int
}
