package accdownload

import(
    "download"
    "config"
    "fmt"
    "code.google.com/p/mahonia"
)

type FiDownloader struct{
    api config.ServiceAPI
    decoder mahonia.Decoder
}

func (d *FiDownloader) Init(){
    const id = "sina-price"
    cm := config.NewServiceConfigManager()
    d.api = cm.GetApi(id, "financialindex")
    d.decoder = mahonia.NewDecoder("gbk")
}

func (d *FiDownloader) GetData(code string) string{
    url := fmt.Sprintf(d.api.Uri, code)
    result := download.HttpGet(url)
    return d.decoder.ConvertString(result)
}

func NewFiDownloader() *FiDownloader{
    d := new(FiDownloader)
    d.Init()

    return d
}
