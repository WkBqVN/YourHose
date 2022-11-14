CREATE TABLE
    PERSONSTATUS (
        PersonID INT NOT NULL IDENTITY(1, 1) PRIMARY KEY,
        Fund BIGINT,
        SleepRoom INT,
        DistrictRequiment VARCHAR(20),
        PriceRange BIGINT,
    );