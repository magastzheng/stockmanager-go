package accdownload_test

import(
    "testing"
    "download/accdownload"
    "fmt"
)

func Test_Account_GetBalanceData(t *testing.T){
    d := accdownload.NewAccountDownloader()
    res := d.GetBalanceData("600001")
    fmt.Println(res)
}

func Test_Account_GetIncomeData(t *testing.T){
    d := accdownload.NewAccountDownloader()
    res := d.GetIncomeData("600001")
    fmt.Println(res)
}

func Test_Account_GetCashflowData(t *testing.T){
    d := accdownload.NewAccountDownloader()
    res := d.GetCashFlowData("600001")
    fmt.Println(res)
}
