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

    key string
}

func (d *SHSEDownloader) Init() {
    d.config = config.NewServiceConfigManager()
    d.key = "shse"
    d.listapi = d.config.GetApi(d.key, "stocklist")
    d.companyapi = d.config.GetApi(d.key, "company")
}

func (d *SHSEDownloader) GetList() string {
    fmt.Println(d.listapi.Uri)
    return HttpGet(d.listapi.Uri)
}

func (d *SHSEDownloader) GetCompanyInfo(code string) string{
    fmt.Println(d.companyapi.Uri)
    
    rand.Seed(time.Now().UnixNano())
    randn := rand.Int63()
    url := fmt.Sprintf(d.companyapi.Uri, randn, randn, code)
    //return HttpGet(url)
    fmt.Println("after:", url)
    header := make(map[string]string)
    header["Referer"] = url
    return HttpGetWithHeader(url, header)
}

func NewSHSEDownloader() *SHSEDownloader{
    d := new(SHSEDownloader)
    d.Init()

    return d
}
