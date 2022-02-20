CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

CREATE TABLE IF NOT EXISTS users(
    ID int auto_increment primary key,
    Name varchar(50) not null,
    Nickname varchar(50) not null unique,
    Email varchar(50) not null unique,
    Password varchar(255) not null,
    CreatedAt timestamp default current_timestamp()
) ENGINE=INNODB;
