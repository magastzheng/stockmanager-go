package dbcreator_test

import(
    "testing"
    "dbcreator"
    //"fmt"
)

func Test_FiDBCreator_Process(t *testing.T){
    m := dbcreator.NewFiDBCreator()
    m.Process()
}
