package models

import (
	// "crud_project/db"
	"crud_project/db"
	"crud_project/logger"
	"crud_project/utility"
	"errors"

	_ "github.com/pelletier/go-toml/query"
)

//
type User struct{
	UserID int64				
	Password string         
	Email string
	JoiningDate string 
}

var user = []User{} // empty slice to save user

func (u User) Save() error{ 
	//Function to add the user data to the database
	query := `INSERT INTO user (email, password, joiningdate)
	VALUES 
	(?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil{
		logger.ErrorLogger.Print("error in sql prepare: error",err)
		return err
	}
	defer stmt.Close()
	password, err := utility.HashPassword(u.Password)
	if err != nil{
		logger.ErrorLogger.Print("error in sql exec: error",err)
		return err
	}
	u.JoiningDate= utility.GetJoining()
	result, err := stmt.Exec(u.Email, password, u.JoiningDate)

	if err != nil{
		logger.ErrorLogger.Print("error in sql exec")
		return err
	}
	id, err := result.LastInsertId()
	if err != nil{
		logger.ErrorLogger.Print("error in lastinsertid", err)
	}
	u.UserID = id
	return err

}

func GetAllUsers() ([]User, error){
	// Read all users details

	query := "select * from user"
	rows, err:= db.DB.Query(query) 

	if err!= nil{
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var u User

		err := rows.Scan(&u.UserID, &u.Email, &u.Password, &u.JoiningDate)

		if err!=nil{
			return nil, err
		}
		user = append(user, u)
	}	
	return user, nil
}

func (u User) ValidateCredentials() error{
	query := "select email, password from users where email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrivedpassword string
	err := row.Scan(&retrivedpassword)
	if err!= nil{
		logger.ErrorLogger.Print("User not found! ", err)
		return err
	}
	isvalid := utility.CheckPassword(u.Password, retrivedpassword)
	
	if !isvalid{
		return errors.New("Creentials invalid")
	} 

	return nil 
}