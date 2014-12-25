package main

import(
    "dbcreator/accgenerator"
)

func main() {
    m := accgenerator.NewAccDBCreator()
    m.Process()
}
