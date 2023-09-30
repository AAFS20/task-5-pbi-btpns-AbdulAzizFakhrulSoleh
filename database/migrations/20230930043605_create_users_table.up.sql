create table users(
    id_u int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(25),
    email VARCHAR(50),
    password TEXT(250),
    created_at DATETIME,
    updated_at DATETIME
);