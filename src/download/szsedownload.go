package download

import(
    "config"
    "code.google.com/p/mahonia"
)

type SZSEDownloader struct {
    config *config.ServiceConfigManager
    listapi config.ServiceAPI
    decoder mahonia.Decoder

    key string
}

func (d *SZSEDownloader) Init() {
    d.config = config.NewServiceConfigManager()
    d.key = "szse"
    d.listapi = d.config.GetApi(d.key, "stocklist")
    d.decoder = mahonia.NewDecoder("gbk")
}

func (d *SZSEDownloader) GetList() string {
    result := HttpGet(d.listapi.Uri)
    return d.decoder.ConvertString(result)
}

func NewSZSEDownloader() *SZSEDownloader{
    d := new(SZSEDownloader)
    d.Init()

    return d
}
