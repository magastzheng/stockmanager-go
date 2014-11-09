package listmanager_test

import(
    "testing"
    "listmanager"
    //"util"
)

func Test_ProcessRealtimeDataManager(t *testing.T) {
    manager := listmanager.NewStockRtDataManager()
    manager.Process()
}

