package dto

import (
	"crud_project/models"
	"errors"
)


// error message template, especifically for missing inputs
func ifEmpty(name , message string) ( error){
	if name == ""{
		err := errors.New(message)  
		return err
	}
	return nil
}



func CheckInput(ev models.Event) (errmap []error){
    err :=  ifEmpty(ev.Name, "name is required")
	if err!= nil{
		errmap = append(errmap, err)
	}
	err =  ifEmpty(ev.Location,"Location is required")
	if err!= nil{
		errmap = append(errmap, err)
	}
	err =  ifEmpty(ev.Description,"Description is required")
	if err!= nil{
		errmap = append(errmap, err)
	}
	return  errmap
}
func CheckUserInput(u models.User) (errmap []error){
    err :=  ifEmpty(u.Email, "Email is required")
	if err!= nil{
		errmap = append(errmap, err)
	}

	err = ifEmpty(u.Password, "Password is required")
	if err!= nil{
		errmap = append(errmap, err)
	}
	if err==nil && len(u.Password) < 8{
		err = errors.New("password must have atleast 8 characters")
		errmap = append(errmap, err)
	}	
	// check length 
	
	return  errmap
}
