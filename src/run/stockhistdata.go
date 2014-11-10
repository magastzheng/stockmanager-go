package main

import(
    "manager"
    //"fmt"
    "util"
)

func main() {
    logger := util.NewLog()
    logger.Info("Start to get historical data")
    m := manager.NewStockHistDataManager()
    m.Process()
    logger.Info("Stock Historical Data complete!")
}
