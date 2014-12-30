package nsmanager_test

import(
    "testing"
    "manager/nsmanager"
    //"fmt"
)

func Test_NSIndexManager_Process(t *testing.T) {
    m := nsmanager.NewNSIndexManager("2014-01", "-1")
    m.Process()
}
