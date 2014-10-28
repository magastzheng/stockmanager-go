package listmanager_test

import (
    "testing"
    "listmanager"
)

func Test_Industry(t *testing.T) {
    manager := new(listmanager.IndustryManager)
    manager.Init()
    manager.Process()
}


