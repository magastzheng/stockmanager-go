package manager_test

import(
    "testing"
    "manager"
)

func Test_NationStatManager_GetData(t *testing.T){
    m := manager.NewNationStatManager()
    m.Process()
}
