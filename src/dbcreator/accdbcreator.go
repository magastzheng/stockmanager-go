package dbcreator

//import(
//
//)

type AccDBCreator struct{
    AccDBCreatorBase
}

func (m *AccDBCreator) Process() {
    m.parser.Parse("../resource/account/accountdb.xlsx")
    dbTabs := ConvertToDBTable(m.parser.CategoryColumnMap) 

    m.CreateDB(dbTabs)
}

func NewAccDBCreator() *AccDBCreator{
    m := new(AccDBCreator)
    m.Init()

    return m
}
