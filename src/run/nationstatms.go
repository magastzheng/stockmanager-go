package main

import(
    "manager"
    "util"
)

func main(){
    util.NewLog().Info("Start to handle money supply")
    m := manager.NewNSMSManager()
    m.Process()
    util.NewLog().Info("Complete to handle money supply")
}
