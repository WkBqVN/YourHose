CREATE TABLE
    HOUSEPRICE (
        HouseID int,
        Timeline DATE,
        Price int,
        FOREIGN KEY (HouseID) REFERENCES HouseInfo(HouseID)
    );