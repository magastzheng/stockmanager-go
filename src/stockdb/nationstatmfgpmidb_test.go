package stockdb_test

import(
    "testing"
    "stockdb"
    "entity"
)

func Test_NSMfgPmiDB_Insert(t *testing.T){
    ms := entity.NSMfgPmi{
        NSPmiCommon:entity.NSPmiCommon{
            "2014-10-31",
            52.12,
            51.23,
            49.42,
            54.45,
            40.2,
            48.21,
            62.1,
        },
        Production: 45.78,
        PurchasingVolume: 47.32,
        Import: 51.23,
        MainRawMaterial: 49.12,
        RawMaterialInventory: 48.52,
        PbizActExpectation: 44.51,
    }

    db := stockdb.NewNSMfgPmiDB("macroindecis")
    db.Delete(ms.Date)
    db.Insert(ms)
}
