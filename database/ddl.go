package database

const createTable = `CREATE TABLE Info
(
    id INT PRIMARY KEY AUTO_INCREMENT,
    curr_to_curr varchar(20) not null unique,
    raw_display text not null
);`
