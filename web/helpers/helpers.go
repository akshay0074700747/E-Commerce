package helpers

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Hash_pass(pass string) (string, error) {

	password := []byte(pass)

	hashedpass, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	return string(hashedpass), err

}

func RandomExpiry() time.Time {

	currentTime := time.Now()

	rand.New(rand.NewSource(time.Now().UnixNano()))

	randomDays := rand.Intn(4) + 3

	expiry := currentTime.Add(time.Duration(randomDays*24) * time.Hour)

	return expiry

}

func VerifyPassword(password, checkpassword string) error {

	return bcrypt.CompareHashAndPassword([]byte(checkpassword), []byte(password))

}
