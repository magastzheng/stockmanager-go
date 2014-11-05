package util_test

import (
    "testing"
    "util"
)

func Test_WriteFile(t *testing.T) {
    util.WriteFile("test.test", "Hello, Test!")
}
