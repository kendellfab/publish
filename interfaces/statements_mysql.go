// +build !sqlite

package interfaces

var CREATE_COMMENT = `CREATE TABLE comment (
id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
page VARCHAR(256) NOT NULL,
username VARCHAR(64),
email VARCHAR(64),
date VARCHAR(64),
content TEXT,
approved TINYINT DEFAULT 0,
day INTEGER, 
month INTEGER, 
year INTEGER)`

var CREATE_USER = `CREATE TABLE user (
id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
name VARCHAR(256) NOT NULL,
email VARCHAR(256) NOT NULL,
password VARCHAR(256) NOT NULL,
role VARCHAR(128) NOT NULL
)`

var CREATE_CATEGORY = `CREATE TABLE category (
id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
title VARCHAR(256) NOT NULL,
slug VARCHAR(256) NOT NULL,
created VARCHAR(128)
)`

var CREATE_POST = `CREATE TABLE post (
id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
title VARCHAR(256) NOT NULL,
slug VARCHAR(256) NOT NULL,
author INTEGER NOT NULL,
created VARCHAR(128) NOT NULL,
content TEXT NOT NULL,
type VARCHAR(128),
published TINYINT,
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
wasRead TINYINT NOT NULL
)`

var CREATE_PAGE = `CREATE TABLE page (
id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
title VARCHAR(256) NOT NULL,
slug VARCHAR(256) NOT NULL,
created VARCHAR(128) NOT NULL,
content TEXT NOT NULL,
published INTEGER
)`

var CREATE_VIEW = `CREATE TABLE view(
id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
who VARCHAR(64) NOT NULL,
at VARCHAR(64) NOT NULL,
type INTEGER,
target VARCHAR(256) NOT NULL)`

var CREATE_RESET = `CREATE TABLE reset(
id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
who INTEGER NOT NULL,
created VARCHAR(128) NOT NULL,
expires VARCHAR(128) NOT NULL,
token VARCHAR(128) NOT NULL);`
