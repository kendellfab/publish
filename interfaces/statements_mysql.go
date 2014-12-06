// +build mysql

package interfaces

var CREATE_COMMENT = `CREATE TABLE comment (
id int PRIMARY KEY AUTO_INCREMENT NOT NULL,
page VARCHAR(256) NOT NULL,
username VARCHAR(64),
email VARCHAR(64),
date VARCHAR(64),
content TEXT,
approved int DEFAULT 0,
day int, 
month int, 
year int)`

var CREATE_USER = `CREATE TABLE user (
id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
name TEXT NOT NULL,
email TEXT NOT NULL,
password TEXT NOT NULL,
role TEXT NOT NULL
)`

var CREATE_CATEGORY = `CREATE TABLE category (
id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
title TEXT NOT NULL,
slug TEXT NOT NULL,
created TEXT
)`

var CREATE_POST = `CREATE TABLE post (
id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
title VARCHAR(256) NOT NULL,
slug VARCHAR(256) NOT NULL,
author INTEGER NOT NULL,
created VARCHAR(128) NOT NULL,
content TEXT NOT NULL,
type VARCHAR(128),
published INTEGER,
tags VARCHAR(128),
category INTEGER,
day INTEGER,
month INTEGER,
year INTEGER
)`

var CREATE_CONTACT = `CREATE TABLE contact (
id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
name VARCHAR(128) NOT NULL,
email VARCHAR(128) NOT NULL,
message TEXT NOT NULL,
read INTEGER NOT NULL)`

var CREATE_PAGE = `CREATE TABLE page (
id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
title VARCHAR(256) NOT NULL,
slug VARCHAR(256) NOT NULL,
created VARCHAR(128) NOT NULL,
content TEXT NOT NULL,
published INTEGER
)`
