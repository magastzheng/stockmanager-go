package listmanager_test

import (
    "testing"
    "listmanager"
)

func Test_Industry(t *testing.T) {
    manager := listmanager.NewIndustryManager()
    manager.Process()
}


