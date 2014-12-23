package stmanager_test

import(
    "testing"
    "manager/stmanager"
)

func Test_SZSECompanyManager_Process(t *testing.T){
    m := stmanager.NewSZSECompanyManager()
    m.Process()
}

