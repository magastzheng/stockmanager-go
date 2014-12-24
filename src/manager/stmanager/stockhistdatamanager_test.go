package stmanager_test

import (
    "testing"
    "manager/stmanager"
)

func Test_StockHistDataManager(t *testing.T) {
    m := stmanager.NewStockHistDataManager()
    m.Process()
}
