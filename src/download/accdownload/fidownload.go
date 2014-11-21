package accdownload

import(
    "download"
    "config"
    "fmt"
)

type FiDownloader struct{
    api config.ServiceAPI
}

func (d *FiDownloader) Init(){
    const id = "sina-price"
    cm := config.NewServiceConfigManager("../../config/serviceconfig.json")
    d.api = cm.GetApi(id, "financialindex")
}

func (d *FiDownloader) GetData(code string) string{
    url := fmt.Sprintf(d.api.Uri, code)
    return download.HttpGet(url)
}

func NewFiDownloader() *FiDownloader{
    d := new(FiDownloader)
    d.Init()

    return d
}
