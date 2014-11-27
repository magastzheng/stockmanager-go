package dbcreator

import(
    acc "entity/accountentity"
    "entity/dbentity"
)

func ConvertToDBTable(tabMap map[string][]*acc.Column) []*dbentity.DBTable {
    comDBCols := make([]*dbentity.DBColumn, 0)
    pcols, ok := tabMap[acc.Common]
    if ok {
		comDBCols = ConvertToDBColumn(pcols, true)
    }
    
    dbTabs := make([]*dbentity.DBTable, 0)
    for k, cols := range tabMap{
        if k == acc.Common {
            continue
        }

        dbTab := dbentity.DBTable{
            TableName: k,
        }
        dbTab.Columns = make([]*dbentity.DBColumn, 0)
        dbTab.Keys = make([]*dbentity.DBColumn, 0)
        dbTab.Columns = append(dbTab.Columns, comDBCols ... )
		dbColumns := ConvertToDBColumn(cols, false)
		dbTab.Columns = append(dbTab.Columns, dbColumns ... )
        dbTabs = append(dbTabs, &dbTab)
    }

    return dbTabs
}

func ConvertToDBColumn(columns []*acc.Column, isnotnull bool) []*dbentity.DBColumn{
	dbColumns := make([]*dbentity.DBColumn, 0)
	for _, col := range columns {
		//don't add duplicated column		
		if !IsExisted(dbColumns, col.Column) {
			dbcol := dbentity.DBColumn{
				Name: col.Column,
				Type: col.Type,
				Maxsize: col.Maxsize,
				IsNotNull: isnotnull,
			}

			dbColumns = append(dbColumns, &dbcol)
		}
	}
	
	return dbColumns
}

func IsExisted(columns []*dbentity.DBColumn, colName string) bool {
	for _, col := range columns {
		if col.Name == colName {
			return true
		}
	}
	
	return false
}
