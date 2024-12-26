package db

import(
	"database/sql"
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
	user_id INTEGER
	)
	`
	_, err := DB.Exec(create_sql)

	if err != nil{
		panic("Could not create table. ")
	}

}
