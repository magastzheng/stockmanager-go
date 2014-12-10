package download

import(
    "fmt"
    "config"
    "time"
    "math/rand"
)

type SHSEDownloader struct {
    config *config.ServiceConfigManager
    listapi config.ServiceAPI
    companyapi config.ServiceAPI
    certapi config.ServiceAPI

    key string
    qkey string
}

func (d *SHSEDownloader) Init() {
    d.config = config.NewServiceConfigManager()
    d.key = "shse"
    d.qkey = "shse-query"
    d.listapi = d.config.GetApi(d.key, "stocklist")
    d.companyapi = d.config.GetApi(d.qkey, "company")
    d.certapi = d.config.GetApi(d.key, "company-cert")
}

func (d *SHSEDownloader) GetList() string {
    fmt.Println(d.listapi.Uri)
    return HttpGet(d.listapi.Uri)
}

func (d *SHSEDownloader) GetCompanyInfo(code string) string{
    fmt.Println(d.companyapi.Uri)
    
    rand.Seed(time.Now().UnixNano())
    randf := rand.Float32()

    randn := int(randf * (100000000+1))

    url := fmt.Sprintf(d.companyapi.Uri, randn, randn, code)
    //return HttpGet(url)
    fmt.Println("after:", url)
    header := make(map[string]string)
    header["Referer"] = d.certapi.Uri
    return HttpGetWithHeader(url, header)
}

func NewSHSEDownloader() *SHSEDownloader{
    d := new(SHSEDownloader)
    d.Init()

    return d
}
