package download_test

import(
    "testing"
    "download"
    "fmt"
    "util"
)

func Test_SHSEDownloader_GetList(t *testing.T) {
    d := download.NewSHSEDownloader()
    res := d.GetList()

    fmt.Println(len(res))
    fmt.Println(res)
}

func Test_SHSEDownloader_GetCompanyInfo(t *testing.T) {
    d := download.NewSHSEDownloader()
    res := d.GetCompanyInfo("600000")

    fmt.Println(len(res))
    fmt.Println(res)
    util.WriteFile("../data/shse600000.dat", res) 
}
