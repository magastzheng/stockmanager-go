package main

import(
    "manager"
)

func main() {
    m := manager.NewNSNonMfgPmiManager()
    m.Process()
}
