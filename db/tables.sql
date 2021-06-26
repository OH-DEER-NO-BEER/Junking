CREATE DATABASE `Junking` DEFAULT CHARACTER SET utf8mb4;
USE `Junking`;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
id INT UNSIGNED NOT NULL AUTO_INCREMENT,
name VARCHAR(16) NOT NULL,
address VARCHAR(256) NOT NULL,
rock_num INT UNSIGNED NOT NULL DEFAULT 0,
scissors_num INT UNSIGNED NOT NULL DEFAULT 0,
paper_num INT UNSIGNED NOT NULL DEFAULT 0,
win_rock_num INT UNSIGNED NOT NULL DEFAULT 0,
win_scissors_num INT UNSIGNED NOT NULL DEFAULT 0,
win_paper_num INT UNSIGNED NOT NULL DEFAULT 0,
PRIMARY KEY(id)
);

-- insert into users(id, name , address, rock_num, scissors_num, paper_num, win_rock_num, win_scissors_num, win_paper_num) values(1, 'hiroto' ,'oshimihiroto@gmail.com', 10, 20, 30, 10, 10, 15);