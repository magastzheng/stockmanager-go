package download

import (
    "config"
    "fmt"
)

const (
    Spring = iota + 1
    Summer
    Autumn
    Winter
)

type StockHistDownloader struct {
    mainSource config.StockHistSourceItem
    seasonSource config.StockHistSourceItem
}

func (s *StockHistDownloader) getMainUrl(code string) string {
    url := fmt.Sprintf(s.mainSource.Url, code)
    return url
}

func (s *StockHistDownloader) getSeasonUrl(code string, year int, season int) string {
    url := fmt.Sprintf(s.seasonSource.Url, code, year, season)
    return url
}

func (s *StockHistDownloader) GetMainPage(code string) string {
    url := s.getMainUrl(code)
    resp := HttpGet(url)

    return resp
}

func (s *StockHistDownloader) GetSeasonPage(code string, year int, season int) string {
    url := s.getSeasonUrl(code, year, season)
    resp := HttpGet(url)

    return resp
}

func NewStockHistDownloader() *StockHistDownloader {
    manager := config.NewStockHistConfig("../config/stockhistconfig.json")
    downloader := new(StockHistDownloader)
    downloader.mainSource = manager.GetConfig("sina")
    downloader.seasonSource = manager.GetConfig("sina-season")

    return downloader
}
