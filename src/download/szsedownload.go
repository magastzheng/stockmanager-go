package download

import(
    "config"
)

type SZSEDownloader struct {
    config *config.ServiceConfigManager
    listapi config.ServiceAPI

    key string
}

func (d *SZSEDownloader) Init() {
    d.config = config.NewServiceConfigManager()
    d.key = "szse"
    d.listapi = d.config.GetApi(d.key, "stocklist")
}

func (d *SZSEDownloader) GetList() string {
    return HttpGet(d.listapi.Uri)
}

func NewSZSEDownloader() *SZSEDownloader{
    d := new(SZSEDownloader)
    d.Init()

    return d
}
