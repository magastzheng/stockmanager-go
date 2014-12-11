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
    fmt.Println(d.listapi.Uri)
    return HttpGet(d.listapi.Uri)
}

func (d *SHSEDownloader) GetCompanyInfo(code string) string{
    fmt.Println(d.companyapi.Uri)
    fmt.Println(d.certapi.Uri)

    rand.Seed(time.Now().UnixNano())
    randf := rand.Float32()

    randn := int(randf * (100000000+1))

    url := fmt.Sprintf(d.companyapi.Uri, randn, randn, code)
    certurl := fmt.Sprintf(d.certapi.Uri, code)
    header := make(map[string]string)
    header["Referer"] = certurl
    //return HttpGet(url)
    fmt.Println("after:", url)
    fmt.Println("Referer:", header, certurl)
    return HttpGetWithHeader(url, header)
}

func (d *SHSEDownloader) GetCompanyIncpt(code string) string{
    fmt.Println(d.compincptapi.Uri)
    fmt.Println(d.certapi.Uri)

    rand.Seed(time.Now().UnixNano())
    randf := rand.Float32()

    randn := int(randf * (100000000+1))

    url := fmt.Sprintf(d.compincptapi.Uri, randn, randn, code)
    certurl := fmt.Sprintf(d.certapi.Uri, code)
    header := make(map[string]string)
    header["Referer"] = certurl
    //return HttpGet(url)
    fmt.Println("after:", url)
    fmt.Println("Referer:", header, certurl)
    return HttpGetWithHeader(url, header)
}

func NewSHSEDownloader() *SHSEDownloader{
    d := new(SHSEDownloader)
    d.Init()

    return d
}
