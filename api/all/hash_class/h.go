package hashpaww

import (
	//"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Mhash(passworsd string) string {
	password := []byte(passworsd)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, 10)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func Vcheck(word, password string, bo chan bool) {

	err := bcrypt.CompareHashAndPassword([]byte(word), []byte(password))

	if err == nil {
		bo <- true
	} else {
		bo <- false
	}
}
func AsyncMhash(passworsd string, channal chan string) {
	password := []byte(passworsd)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, 10)
	if err != nil {
		panic(err)
	}
	channal <- string(hashedPassword)

}
