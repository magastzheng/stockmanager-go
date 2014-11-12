use macroindecis;

drop table if exists moneysupply;

create table moneysupply(
    date date not null primary key,
    m0 decimal(15,2) default null,
    m0pct float,
    m1 decimal(15,2) default null,
    m1pct float,
    m2 decimal(15,2) default null,
    m2pct float
);

drop table if exists mfgpmi;

create table mfgpmi(
    date date not null primary key,
    pmi float,
    prodution float,
    neworder float,
    newexportorder float,
    inhandorder float,
    inventory float,
    purchasingvolume float,
    import float,
    mainrawmaterialpurchaseprice float,
    rawmaterialsinventory float,
    employees float,
    supplierdeliverytime float,
    pbizactexpectation float
);

drop table if exists nonmfgpmi;

create table nonmfgpmi(
    date date not null primary key,
    pmi float,
    neworder float,
    newexportorder float,
    inhandorder float,
    inventory float,
    iminputprice float,
    subscriptionprice float,
    emploees float,
    supplierdeliverytime float,
    bizactexpectation float
);

