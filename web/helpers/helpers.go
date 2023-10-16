package helpers

import "golang.org/x/crypto/bcrypt"

func Hash_pass(pass string) string {

	password := []byte(pass)

	hashedpass, err := bcrypt.GenerateFromPassword(password,bcrypt.DefaultCost)

	if err != nil {
		panic("Password cannot be hashed...")
	}

	return string(hashedpass)
	
}