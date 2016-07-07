# README

## Database Design

1. Create database user dbuser

	`CREATE USER dbuser WITH PASSWORD 'dogtutu';`

2. Create database ownered by dbuser

	`CREATE DATABASE tantan_db OWNER dbuser;`

3.  Grant all privileges on database tantan_db

	`GRANT ALL PRIVILEGES ON DATABASE tantan_db to dbuser;`

4. Create following tables

	- User table

        |field|type|description|
        |-----|----|-----------|
        |id|varchar(64)|user id|
        |name|varchar(32)|user name|
        |type|varchar(32)|data type|
        
        ```

        CREATE TABLE user_tbl(
            id VARCHAR(64) NOT NULL DEFAULT '', 
            name VARCHAR(32) NOT NULL DEFAULT '',
            type VARCHAR(32) NOT NULL DEFAULT '',
            PRIMARY KEY (id)
        );
        ```
        
	- Relationship table

        |field|type|description|
        |-----|----|-----------|
        |user_id|varchar(64)|user id|
        |peer_user_id|varchar(64)|peer user id|
        |state|varchar(32)|user relationship|
        |type|varchar(32)|data type|

        ```
        CREATE TABLE relationship_tbl(
            user_id VARCHAR(64) NOT NULL DEFAULT '',
            peer_user_id VARCHAR(64) NOT NULL DEFAULT '', 
            state VARCHAR(32) NOT NULL DEFAULT '',
            type VARCHAR(32) NOT NULL DEFAULT '',
            PRIMARY KEY (user_id, peer_user_id)
        );
        ```

## How to Run demo

1. Install PostgreSQL database(>= 9.5.1), and add database/tables referring to the above steps.

2. Install Go development environment(I use the version 1.5, you can choose yourself version, just be nice to use).

3. Install dependent Go packages:
	- pg
	
		`go get gopkg.in/pg.v4`

	- mux
		
		`go get github.com/gorilla/mux`
	
4. Run the demo
	- Append the project path to GOPATH env
	
		`export GOPATH=$GOPATH:path_to_project`
		
	- Build and install the demo:
		
		`go install tantan-server`
			
	- Run the demo:
		
		`path_to_project/bin/tantan-server`
		