CREATE TABLE
    HOUSEDETAIL (
        HouseDetailID INT NOT NULL IDENTITY(1, 1),
        HouseID int,
        HouseIntro text NOT Null,
        HouseComment text NOT NULL,
        PRIMARY KEY (HouseID),
        FOREIGN KEY (HouseID) REFERENCES HouseInfo(HouseID)
    );