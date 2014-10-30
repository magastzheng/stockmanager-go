package util_test

import (
    "testing"
    "util"
)

func Test_IsStringNotEmpty(t *testing.T){
    res := util.IsStringNotEmpty("")
    if res {
        t.Log("success")
    }
}

func Test_WriteFile(t *testing.T) {
    util.WriteFile("test.test", "Hello, Test!")
}
