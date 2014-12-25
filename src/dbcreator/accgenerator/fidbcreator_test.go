package accgenerator_test

import(
    "testing"
    "dbcreator/accgenerator"
    //"fmt"
)

func Test_FiDBCreator_Process(t *testing.T){
    m := accgenerator.NewFiDBCreator()
    m.Process()
}

