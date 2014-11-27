package accmanager_test

import(
    "testing"
    "manager/accmanager"
    //"fmt"
)

func Test_AccountManager_Process(t *testing.T){
    m := accmanager.NewAccountManager()
    m.Process()
}

