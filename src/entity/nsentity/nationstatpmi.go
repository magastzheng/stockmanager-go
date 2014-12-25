package nsentity

type NSPmiCommon struct{
    Date string
    Pmi float32
    NewOrder float32
    NewExportOrder float32
    InHandOrder float32
    Inventory float32
    Employees float32
    SupplierDeliveryTime float32
}

type NSMfgPmi struct{
    NSPmiCommon
    Production float32
    PurchasingVolume float32
    Import float32
    MainRawMaterial float32
    RawMaterialInventory float32
    PbizActExpectation float32
}

type NSNonMfgPmi struct{
    NSPmiCommon
    ImInputPrice float32
    SubscriptionPrice float32
    BizActExpectation float32
}
