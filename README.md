# README

## Database Design

### tantan_db database

#### user table

|field|type|description|
|-----|----|-----------|
|id|varchar(64)|user id|
|name|varchar(32)|user name|
|type|varchar(32)|data type|


#### relationship table

|field|type|description|
|-----|----|-----------|
|id|varchar(64)|user id|
|user_id|varchar(64)|related user id|
|state|varchar(32)|user relationship|
|type|varchar(32)|data type|

