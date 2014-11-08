package main

import(
    "listmanager"
    //"fmt"
    "util"
)

func main() {
    logger := util.NewLog()
    logger.Info("Start to get historical data")
    manager := listmanager.NewStockHistDataManager()
    manager.Process()
    logger.Info("Stock Historical Data complete!")
}
