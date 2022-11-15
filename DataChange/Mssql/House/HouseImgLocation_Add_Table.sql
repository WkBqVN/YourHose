CREATE TABLE
    HOUSEIMG (
        ImageID int NOT NULL IDENTITY(1, 1),
        LinkLocal1 text NOT NULL,
        LinkLocal2 text,
        LinkLocal3 text,
        HouseID int,
        PRIMARY KEY (ImageID),
        FOREIGN KEY (HouseID) REFERENCES HouseInfo(HouseID)
    );