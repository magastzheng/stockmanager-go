package entity

type Stock struct{
    Id string
    Name string
    Exchange string
}

type StockSummary struct{
    Stock
    Website string
}
