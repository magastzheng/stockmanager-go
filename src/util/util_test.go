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


