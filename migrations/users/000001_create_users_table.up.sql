CREATE TABLE IF NOT EXISTS users(
    Email varchar(255) NOT NULL PRIMARY KEY,
    Password varchar(255) NOT NULL,
    Name varchar(255) NOT NULL,
    IsAdmin int NOT NULL
);
