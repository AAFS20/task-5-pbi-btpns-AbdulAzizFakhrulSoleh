create table photos(
    idp int PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(250),
    caption TEXT,
    photo_url TEXT,
    user_id int NOT NULL,
    FOREIGN KEY(user_id)
        REFERENCES users (id_u)
        ON DELETE CASCADE
);