package encryption

import (
	"log"
	"golang.org/x/crypto/bcrypt"
)

// @reference https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
func HashAndSalt(pwd []byte) string {
    // MinCost is just an integer constant provided by the bcrypt
    // package along with DefaultCost & MaxCost. 
    // The cost can be any value you want provided it isn't lower
    // than the MinCost (4)
	// return password yang dienkripsi 
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println("Error hash and salt <package encryption> : ", err.Error())
	}

	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
    // will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)

	if err := bcrypt.CompareHashAndPassword(byteHash, plainPwd); err != nil {
		log.Println("Error when compare hash and password <package encryption> : ", err.Error())
		return false 
	}

	return true
}
