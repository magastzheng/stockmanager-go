package dbcreator

import(
    acc "entity/accountentity"
    "entity/dbentity"
)

func ConvertToDBColumn(tabMap map[string][]*acc.Column) []*dbentity.DBTable {
    comDBCols := make([]*dbentity.DBColumn, 0)
    pcols, ok := tabMap[acc.Common]
    if ok {
        for _, col := range pcols {
            dbcol := dbentity.DBColumn{
                Name: col.Column,
                Type: col.Type,
                Maxsize: col.Maxsize,
                IsNotNull: true,
            }

            comDBCols = append(comDBCols, &dbcol)
        }
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

        for _, col := range cols {
            dbcol := dbentity.DBColumn{
                Name: col.Column,
                Type: col.Type,
                Maxsize: col.Maxsize,
            }

            dbTab.Columns = append(dbTab.Columns, &dbcol)
        }

        dbTabs = append(dbTabs, &dbTab)
    }

    return dbTabs
}
