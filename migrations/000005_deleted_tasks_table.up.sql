CREATE TABLE IF NOT EXISTS Deleted_Tasks (
    ID int PRIMARY KEY AUTO_INCREMENT,
    Name varchar(255),
    Detail varchar(512),
    Done boolean,
    Owner varchar(255)
);
