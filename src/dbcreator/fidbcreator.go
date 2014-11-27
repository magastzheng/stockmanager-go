package dbcreator

//import(
//
//)

type FiDBCreator struct{
    AccDBCreatorBase
}

func (m *FiDBCreator) Process() {
    m.parser.Parse("../resource/account/financialindexdb.xlsx")
    dbTabs := ConvertToDBTable(m.parser.CategoryColumnMap) 

    m.CreateDB(dbTabs)
}

func NewFiDBCreator() *FiDBCreator{
    m := new(FiDBCreator)
    m.Init()

    return m
}
