package nsdownload

import(
    "strings"
    "fmt"
    "config"
    "download"
    "time"
)

type NationStatDownloader struct{
    root config.ServiceAPI
    child1 config.ServiceAPI
    child2 config.ServiceAPI
    child3 config.ServiceAPI
    period config.ServiceAPI
    data config.ServiceAPI
}

func (d *NationStatDownloader) Init() {
    const id = "nationstat"
    cm := config.NewServiceConfigManager()
    //cm := config.NewServiceConfigManager1()
    d.root = cm.GetApi(id, "indexroot")
    d.child1 = cm.GetApi(id, "children1")
    d.child2 = cm.GetApi(id, "children2")
    d.child3 = cm.GetApi(id, "children3")
    d.period = cm.GetApi(id, "timeperiod")
    d.data = cm.GetApi(id, "data")
}

func (d *NationStatDownloader) GetServiceData(api config.ServiceAPI, v string) string {
    var result string
    switch api.Method {
        case "GET":
            url := fmt.Sprintf(api.Uri, v)
            result = download.HttpGet(url)
        case "POST":
            query := fmt.Sprintf(api.Data, v)
            result = download.HttpPostForm(api.Uri, query)
    }

    return result
}

func (d *NationStatDownloader) GetRoot() string {
    uri := fmt.Sprintf(d.root.Uri, d.getRand())
    return download.HttpGet(uri)
}

func (d *NationStatDownloader) GetChild(code string, level int) string {
    var result string
    //querycode := strings.Join(codes, ",")
    switch level {
        case 1:
            result = d.GetServiceData(d.child1, code)
        case 2:
            result = d.GetServiceData(d.child2, code)
        case 3:
            result = d.GetServiceData(d.child3, code)
    }

    return result
}

func (d *NationStatDownloader) GetPeriod() string {
    uri := fmt.Sprintf(d.period.Uri, d.getRand())
    return download.HttpGet(uri)
}

func (d *NationStatDownloader) GetData(codes []string, start string, end string) string {
    //start, end should be yyyyMM. The period like '-1,200101' means start from 2001-01 to current
    querycode := strings.Join(codes, ",")

    randn := d.getRand()
    var period string
    if end == "-1" {
        period = fmt.Sprintf("%v,%v", end, start)
    } else {
        period = fmt.Sprintf("%v,%v", start, end)
    }
    
    url := fmt.Sprintf(d.data.Uri, randn, querycode, period)
    return download.HttpGet(url)
}

func (d *NationStatDownloader) getRand() int {
    nanos := time.Now().UnixNano()
    return int(nanos / 1000000)
}

func NewNationStatDownloader() *NationStatDownloader {
    d := new(NationStatDownloader)
    d.Init()
    
    return d
}
