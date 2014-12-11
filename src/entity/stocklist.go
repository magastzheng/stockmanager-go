package entity

type StockItem struct{
    Id string
    Name string
}

type Stock struct{
    StockItem
    Website string
}
