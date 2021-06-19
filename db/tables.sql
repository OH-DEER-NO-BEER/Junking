CREATE DATABASE `Junking` DEFAULT CHARACTER SET utf8mb4;
USE `Junking`;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
id INT UNSIGNED NOT NULL AUTO_INCREMENT,
-- name VARCHAR(16) NOT NULL,
address VARCHAR(256) NOT NULL,
rock_num INT UNSIGNED NOT NULL,
scissors_num INT UNSIGNED NOT NULL,
paper_num INT UNSIGNED NOT NULL,
win_rock_num INT UNSIGNED NOT NULL,
win_scissors_num INT UNSIGNED NOT NULL,
win_paper_num INT UNSIGNED NOT NULL,
PRIMARY KEY(id)
);

insert into users(id, address, rock_num, scissors_num, paper_num, win_rock_num, win_scissors_num, win_paper_num) values(1, 'oshimihiroto@gmail.com', 10, 20, 30, 10, 10, 15);