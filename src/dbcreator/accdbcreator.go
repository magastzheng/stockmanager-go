package dbcreator

//import(
//
//)

type AccDBCreator struct{
    AccDBCreatorBase
}

func (m *AccDBCreator) Process() {
    m.parser.Parse("../resource/account/accountdb.xlsx")
	
	for sheet, category := range m.parser.CategoryColumnMap{
		m.logger.Info("Start to create table for the sheet:", sheet)
		dbTabs := ConvertToDBTable(category) 
		m.CreateDB(dbTabs)
	}
}

func (m *AccDBCreator) Delete(){
	m.parser.Parse("../resource/account/accountdb.xlsx")
	
	for sheet, category := range m.parser.CategoryColumnMap{
		m.logger.Info("Start to create table for the sheet:", sheet)
		dbTabs := ConvertToDBTable(category) 
		m.DropDB(dbTabs)
	}
}

func NewAccDBCreator() *AccDBCreator{
    m := new(AccDBCreator)
    m.Init()

    return m
}
