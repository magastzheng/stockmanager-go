package manager_test

import(
    "testing"
    "manager"
    //"stockdb"
    //"fmt"
)

func Test_NSNonMfgPmiManager_Process(t *testing.T){
    m := manager.NewNSNonMfgPmiManager()
    m.Process()
}
