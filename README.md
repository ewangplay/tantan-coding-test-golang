# README

## Database Design

### tantan_db database

1. Create database user dbuser
CREATE USER dbuser WITH PASSWORD 'password';

2. Create database ownered by dbuser
CREATE DATABASE tantan_db OWNER dbuser;

3.  Grant all privileges on database tantan_db
GRANT ALL PRIVILEGES ON DATABASE tantan_db to dbuser;

4. Login the database tantan_db
psql -U dbuser -d tantan_db -h 127.0.0.1 -p 5432

5. Create following tables

```
** user table **

|field|type|description|
|-----|----|-----------|
|id|varchar(64)|user id|
|name|varchar(32)|user name|
|type|varchar(32)|data type|

CREATE TABLE user_tbl(
    id VARCHAR(64), 
    name VARCHAR(32),
    type VARCHAR(32)
);
```

```
** relationship table **

|field|type|description|
|-----|----|-----------|
|id|varchar(64)|user id|
|user_id|varchar(64)|related user id|
|state|varchar(32)|user relationship|
|type|varchar(32)|data type|

CREATE TABLE relationship_tbl(
    id VARCHAR(64),
    user_id VARCHAR(64), 
    state VARCHAR(32),
    type VARCHAR(32)
);
```

## How to Demonstrate the program

