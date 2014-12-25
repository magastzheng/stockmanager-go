package accgenerator_test

import(
    "testing"
    "dbcreator"
    //"fmt"
)

func Test_FiDBCreator_Process(t *testing.T){
    m := accgenerator.NewAccDBCreator()
    //m.Process()
	m.Delete()
}

