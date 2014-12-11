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
    compincptapi config.ServiceAPI
    certapi config.ServiceAPI

    key string
    qkey string
}

func (d *SHSEDownloader) Init() {
    d.config = config.NewServiceConfigManager()
    d.key = "shse"
    d.qkey = "shse-query"
    d.listapi = d.config.GetApi(d.key, "stocklist")
    d.certapi = d.config.GetApi(d.key, "company-cert")
    //d.companyapi = d.config.GetApi(d.qkey, "company")
    d.companyapi = d.config.GetApi(d.qkey, "company-info")
    d.compincptapi = d.config.GetApi(d.qkey, "compincpt")
}

func (d *SHSEDownloader) GetList() string {
    return HttpGet(d.listapi.Uri)
}

func (d *SHSEDownloader) getData(code string, urifmt string) string {
    rand.Seed(time.Now().UnixNano())
    randf := rand.Float32()

    randn := int(randf * (100000000+1))
    url := fmt.Sprintf(urifmt, randn, randn, code)
    certurl := fmt.Sprintf(d.certapi.Uri, code)
    header := make(map[string]string)
    header["Referer"] = certurl
    return HttpGetWithHeader(url, header)
}

func (d *SHSEDownloader) GetCompanyInfo(code string) string{
    return d.getData(code, d.companyapi.Uri)
}

func (d *SHSEDownloader) GetCompanyIncpt(code string) string{
    return d.getData(code, d.compincptapi.Uri)
}

func NewSHSEDownloader() *SHSEDownloader{
    d := new(SHSEDownloader)
    d.Init()

    return d
}
