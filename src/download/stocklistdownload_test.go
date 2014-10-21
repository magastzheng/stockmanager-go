package download

import (
    "testing"
    "fmt"
)

//http://app.finance.ifeng.com/hq/list.php
//http://app.finance.ifeng.com/hq/list.php?type=stock_a&class=ha
//class can be: ha, sa, gem

func Test_GetPage(t *testing.T) {
    s := new(StockDownloader)
    str := s.GetPage("http://app.finance.ifeng.com/hq/list.php", "stock_a", "sa")

    fmt.Println(len(str))
}

