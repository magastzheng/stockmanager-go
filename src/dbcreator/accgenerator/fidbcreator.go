package accgenerator

//import(
//
//)

type FiDBCreator struct{
    AccDBCreatorBase
}

func (m *FiDBCreator) Process() {
    m.parser.Parse("../../resource/account/financialindexdb.xlsx")
	category := m.parser.GetSheetMap("findex")
    dbTabs := ConvertToDBTable(category) 

    m.CreateDB(dbTabs)
}

func NewFiDBCreator() *FiDBCreator{
    m := new(FiDBCreator)
    m.Init()

    return m
}
