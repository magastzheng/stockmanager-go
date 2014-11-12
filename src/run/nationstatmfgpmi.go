package main

import(
    "manager"
)

func main(){
    m := manager.NewNSMfgPmiManager()
    m.Process()
}
