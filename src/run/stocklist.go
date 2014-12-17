package main

import (
    "manager"
    "manager/stmanager"
    "flag"
    "fmt"
)

func main() {
    t := flag.Int("t", 0, "stock list type")
    flag.Parse()
    
    switch (*t) {
        case 0:
            fmt.Println("Get all the stock list")
            m := manager.NewStockListManager()
            m.Process()
            fmt.Println("Stock list run complete!")
        case 1:
            fmt.Println("Get new stock list")
            shm := stmanager.NewSHSEListManager()
            shm.Process()
            szm := stmanager.NewSZSEListManager()
            szm.Process()
            fmt.Println("New stock list run complete!")
    }
    
}
