package  main

import(
    "manager"
    "manager/nsmanager"
    "fmt"
    "time"
    "flag"
)

func main(){
    start := time.Now()
    
    t := flag.String("t", "ns", "nation stat data type")
    flag.Parse()

    var m manager.Manager
    
    switch *t {
        case "ns":
            m = nsmanager.NewNationStatManager()
        case "pmi":
            m = nsmanager.NewNSMfgPmiManager()
        case "npmi":
            m = nsmanager.NewNSNonMfgPmiManager()
        case "ms":
            m = nsmanager.NewNSMSManager()
    }
    
    m.Process()

    duration := time.Since(start)
    fmt.Printf("Complete to process: %v s\n", duration.Seconds())
}
