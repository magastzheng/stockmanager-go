package download_test

import(
    "testing"
    "fmt"
    "download"
)

func Test_SZSEDownloader_GetList(t *testing.T){
    d := download.NewSZSEDownloader()
    res := d.GetList()
    fmt.Println(len(res))
    fmt.Println(res)
}
