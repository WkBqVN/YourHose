CREATE TABLE
    PERSONINFO (
        PersonID int NOT NULL IDENTITY(1, 1) PRIMARY KEY,
        LastName varchar(20) NOT NULL,
        FirstName varchar(20) NOT NULL,
        PersonalCardID bigint check(
            PersonalCardID between 0000000000000000 and 9999999999999999
        ) NOT NULL,
        BirthDay DATE NOT NULL,
        Gender varchar(4) NOT NULL,
    );