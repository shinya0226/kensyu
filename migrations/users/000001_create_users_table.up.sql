CREATE TABLE IF NOT EXISTS users(
    Email varchar(255) NOT NULL,
    Password varchar(255) NOT NULL,
    Name varchar(255) NOT NULL,
    IsAdmin int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
