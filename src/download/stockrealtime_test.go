package download_test

import(
    "testing"
    "download"
    "fmt"
)

func Test_GetRealData(t *testing.T){
    downloader := download.NewStockRtDownloader()
    res := downloader.GetData("000150", "EX$$$$XSHE")
    fmt.Println(res)
    fmt.Println("====================")
    res = downloader.GetData("601600", "EX$$$$XSHG")
    fmt.Println(res)
}
