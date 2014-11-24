package download

import(
    "config"
    "fmt"
)

type StockRtDownloader struct {
    config config.ServiceAPI
}

func (s *StockRtDownloader) GetData(code, exchange string) string {
    var excode string
    switch exchange {
        case "EX$$$$XSHG":
            excode = "sh"
        case "EX$$$$XSHE":
            excode = "sz"
    }

    url := fmt.Sprintf(s.config.Uri, excode, code)
    resp := HttpGet(url)

    return resp
}

func NewStockRtDownloader() *StockRtDownloader{
    cm := config.NewServiceConfigManager()
    downloader := new(StockRtDownloader)
    downloader.config = cm.GetApi("sina-realtime", "realtime")

    return downloader
}
