package download_test

import (
    "testing"
    "download"
    //"fmt"
)

func Test_GetMainPage(t *testing.T) {
    downloader := download.NewStockHistDownloader()
    resp := downloader.GetMainPage("601390")

    if len(resp) > 0 {
        t.Log(resp)
        t.Log("success")
    }
}

func Test_GetSeasonPage(t *testing.T) {
    downloader := download.NewStockHistDownloader()
    resp := downloader.GetSeasonPage("601390", 2013, 2)

    if len(resp) > 0 {
        t.Log(resp)
        t.Log("success")
    }
}
