package main

import(
    "listmanager"
    "fmt"
)

func main() {
    manager := listmanager.NewStockHistDataManager()
    manager.Process()
    fmt.Println("Stock Historical Data complete!")
}
