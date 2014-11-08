package util_test

import(
    "util"
    "testing"
   // "fmt"
)

func Test_LogError(t *testing.T){
    logger := util.NewLog()
    logger.Error("This is a logging error!")
    logger.Error("This is second logging error!")
}



