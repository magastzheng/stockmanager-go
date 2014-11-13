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
    production float,
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
    employees float,
    supplierdeliverytime float,
    bizactexpectation float
);

create table currentcpi(
    cpi float,
    food float,
    tobaccoliquor float,
    clothing float
    housefacility float,
    health float,
    transport float,
    recreationedu float,
    residence float
);

create table currenturbancpi(
    cpi float,
    food float,
    tobaccoliquor float,
    clothing float
    housefacility float,
    health float,
    transport float,
    recreationedu float,
    residence float
);

create table currentruralcpi(
    cpi float,
    food float,
    tobaccoliquor float,
    clothing float
    housefacility float,
    health float,
    transport float,
    recreationedu float,
    residence float
);


create table foodcpi(
    grain float,
    grease float,
    meat float,
    egg float,
    aquatic float,
    vegetable float,
    fruit float/*,
    milk float*/
);
