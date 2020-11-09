CREATE TABLE IF NOT EXISTS Tokens (
    ID int PRIMARY KEY AUTO_INCREMENT,
    User varchar(255),
    Access_Token varchar(512)
);
