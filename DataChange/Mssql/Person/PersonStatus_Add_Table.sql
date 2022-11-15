CREATE TABLE
    PERSONSTATUS (
        PersonID INT,
        Fund BIGINT,
        SleepRoom INT,
        DistrictRequiment VARCHAR(20),
        PriceRange BIGINT,
        FOREIGN KEY (PersonID) REFERENCES PersonInfo(PersonID)
    );