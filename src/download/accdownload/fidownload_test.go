package accdownload_test

import(
    "testing"
    "download/accdownload"
    "fmt"
)

func Test_FiDownloader(t *testing.T) {
    d := accdownload.NewFiDownloader()
    res := d.GetData("000001")

    fmt.Println(res)
}
