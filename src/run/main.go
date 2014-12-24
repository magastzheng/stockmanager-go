package main

import(
    "manager"
    "manager/stmanager"
    "manager/accmanager"
    "fmt"
    "time"
    "flag"
)

func Call(m manager.Manager) {
    m.Process()
}

func main() {
    var m manager.Manager
    start := time.Now()

    t := flag.String("t", "list", "stock manager type")
    flag.Parse()

    fmt.Println(*t, t)
    switch *t {
        case "list": //get stock list
            m = manager.NewStockListManager()
        case "nlist": //get new stock list
            m = stmanager.NewListManager()
        case "compdata": //get company data
            m = stmanager.NewCompanyManager()
        case "strtdata": //get stock daily data
            m = manager.NewStockRtDataManager()
        case "sthdata": //get stock historical data
            m = manager.NewStockHistDataManager()
        case "nsthdata": //get new stock historical data
            m = stmanager.NewStockHistDataManager()
        case "accdata": //get account data
            m = accmanager.NewAccountManager()
    }

    Call(m)
    duration := time.Since(start)
    fmt.Printf("Complete to process: %v s\n", duration.Seconds())
}
