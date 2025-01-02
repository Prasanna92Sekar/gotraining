package utility

import (
	"time"
	"golang.org/x/crypto/bcrypt"
)

func GetJoining() (string){

	curr_time := time.Now()
	Readable_time := curr_time.Format(("2006-January-02"))
	return Readable_time
}

func HashPassword(password string) (string, error){
	hashbytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashbytes), err
}

func CheckPassword(password, hashpassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password))

	return err == nil 
}
