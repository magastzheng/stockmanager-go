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

    util.WriteFile("../data/shse_list.dat", res)
}

func Test_SHSEDownloader_GetCompanyInfo(t *testing.T) {
    d := download.NewSHSEDownloader()
    res := d.GetCompanyInfo("600000")

    fmt.Println(len(res))
    util.WriteFile("../data/shse600000.dat", res) 
}

func Test_SHSEDownloader_GetCompanyIncpt(t *testing.T){
    d := download.NewSHSEDownloader()
    res := d.GetCompanyIncpt("600000")

    fmt.Println(res)
    util.WriteFile("../data/shse600000_incpt.dat", res) 
}
