package listmanager_test

import(
    "testing"
    "listmanager"
    //"fmt"
)

func Test_Process(t *testing.T) {
    //s := new(StockListManager)
    //s.Init("../config/stocklist.json")
    s := listmanager.NewStockListManager()
    s.Process()
}
