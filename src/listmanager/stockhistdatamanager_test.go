package listmanager_test

import (
    "testing"
    "listmanager"
)

func Test_StockHistDataManager(t *testing.T) {
    manager := listmanager.NewStockHistDataManager()
    manager.Process()
}
