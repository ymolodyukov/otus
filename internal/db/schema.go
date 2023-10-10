package db

var Schema = `
CREATE TABLE IF NOT EXISTS person (
    id varchar(36), 
	firstname  varchar(50),
	lastname varchar(50),
	age        integer,
	sex        varchar(1),
	biography  text ,
	city       varchar(100)
);

CREATE TABLE IF NOT EXISTS credentials (
    id varchar(36),
    user_id varchar(36),
    password varchar(100)
);`
