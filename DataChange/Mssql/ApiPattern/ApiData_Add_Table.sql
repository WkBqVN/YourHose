CREATE TABLE
    APIPATTERN (
        AID INT NOT NULL IDENTITY(1, 1),
        APIID int,
        APIPAYLOAD text,
        APIKEYS text,
        PRIMARY KEY (AID),
    );