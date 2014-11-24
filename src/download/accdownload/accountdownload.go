package accdownload

import(
    "download"
    "config"
    "code.google.com/p/mahonia"
    "fmt"
)

type AccountDownloader struct{
    balance config.ServiceAPI
    income config.ServiceAPI
    cashflow config.ServiceAPI
    decoder mahonia.Decoder
}

func (d *AccountDownloader) Init(){
    const id = "sina-price"
    cm := config.NewServiceConfigManager()
    d.balance = cm.GetApi(id, "balancesheet")
    d.income = cm.GetApi(id, "income")
    d.cashflow = cm.GetApi(id, "cashflow")
    d.decoder = mahonia.NewDecoder("gbk")
}

func (d *AccountDownloader) GetData(url string) string{
    result := download.HttpGet(url)
    return d.decoder.ConvertString(result)
}

func (d *AccountDownloader) GetBalanceData(code string) string{
    url := fmt.Sprintf(d.balance.Uri, code)
    return d.GetData(url)
}

func (d *AccountDownloader) GetIncomeData(code string) string{
    url := fmt.Sprintf(d.income.Uri, code)
    return d.GetData(url)
}

func (d *AccountDownloader) GetCashFlowData(code string) string{
    url := fmt.Sprintf(d.cashflow.Uri, code)
    return d.GetData(url)
}

func NewAccountDownloader() *AccountDownloader{
    d := new(AccountDownloader)
    d.Init()

    return d
}
