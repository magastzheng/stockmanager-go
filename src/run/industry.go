package main

import (
    "manager"
    "manager/industrymanager"
    "flag"
    "time"
    "fmt"
)

func main() {
    t := flag.String("t", "csrc", "Industry information type")
    flag.Parse()
    
    start := time.Now()
    var m manager.Manager
    switch *t {
        case "csrc":
            m = industrymanager.NewIndustryManager()
    }

    m.Process()
    
    duration := time.Since(start)

    fmt.Println("Industry complete!")
    fmt.Printf("Complete to process: %v s\n", duration.Seconds())
}


