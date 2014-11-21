package accmanager_test

import(
    "testing"
    "manager/accmanager"
    //"fmt"
)

func Test_FiManager_Process(t *testing.T){
    m := accmanager.NewFiManager()
    m.Process()
}
