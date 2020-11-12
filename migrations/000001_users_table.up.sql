CREATE TABLE IF NOT EXISTS Users (
    ID int PRIMARY KEY AUTO_INCREMENT,
    Username varchar(255),
    Email varchar(255),
    Password varchar(255),
    Deleted boolean
);
