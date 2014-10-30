package download

//http://app.finance.ifeng.com/hq/list.php
//http://app.finance.ifeng.com/hq/list.php?type=stock_a&class=ha
//class can be: ha, sa, gem

type StockDownloader struct{
	id string
}

func (s *StockDownloader) GetUrl(baseUrl, typ, class string) string {
    return  baseUrl+"?type="+typ+"&class="+class
}

func (s *StockDownloader) GetPage(baseUrl, typ, class string) string {
    url := s.GetUrl(baseUrl, typ, class)
    return HttpGet(url)
}

func NewDownloader() *StockDownloader {
    return &StockDownloader{}
}
