package nsmanager_test

import(
    "testing"
    "manager/nsmanager"
    //"stockdb"
    //"fmt"
)

func Test_NSNonMfgPmiManager_Process(t *testing.T){
    m := nsmanager.NewNSNonMfgPmiManager()
    m.Process()
}
