package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // _ so that it i not removed or give as it will not used directly but indirectly by databse sql. so we use _
)

var DB *sql.DB
 
func InitDB() {
	// Initilize database
    var err error
    DB, err = sql.Open("sqlite3", "api.db")
 
    if err != nil {
        panic("Could not connect to database.")
    }
 
    DB.SetMaxOpenConns(10)
    DB.SetMaxIdleConns(5)
 
	createUserTable()
    createTables()
	
}

func createTables() {
	//Function to create table
	create_sql :=`
	CREATE TABLE IF NOT EXISTS events(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	dateTime DATETIME  NOT NULL,
	user_id INTEGER, 
	FOREIGN KEY(user_id) REFERENCES user(userid)
	)
	`
	_, err := DB.Exec(create_sql)

	if err != nil{
		fmt.Println(err)
		panic("Could not create table. ")
	}

}

func createUserTable(){
	// Function to create user table
	createTable := `CREATE TABLE IF NOT EXISTS user (
	userid INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	joiningdate Text  NOT NULL
	)
	`
	_, err := DB.Exec(createTable)
	if err != nil{
		panic("Could not create User table. ")
	}
}
