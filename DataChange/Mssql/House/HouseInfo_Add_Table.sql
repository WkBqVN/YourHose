CREATE TABLE
    HOUSEINFO (
        HouseID int NOT NULL IDENTITY(1, 1) PRIMARY KEY,
        Country varchar(20) NOT NULL,
        City varchar(20) NOT NULL,
        District varchar(20) NOT NULL,
        Street varchar(20) NOT NULL,
        StreetNumber varchar(20) NOT NULL,
    );