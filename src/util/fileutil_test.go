package util_test

import (
    "testing"
    "util"
    "os"
)

func Test_WriteFile(t *testing.T) {
    util.WriteFile("test.test", "Hello, Test!")
}

func Test_WriteFileToFolder(t *testing.T) {
    folder := "../testfolder/test/"
    filename := folder + "test.dat"
    util.WriteFile(filename, "test data into testing")
    os.RemoveAll(filename)
    os.Remove("../testfolder")
}
