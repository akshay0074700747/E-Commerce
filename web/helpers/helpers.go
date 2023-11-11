package helpers

import (
	"fmt"
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

func StatusCheck(code int) (bool, string) {

	var statuscheck = map[int]string{
		0: "processing",
		1: "shipped",
		2: "delayed delivery",
		3: "out for delivery",
		4: "delivered",
	}

	status, exists := statuscheck[code]

	return exists, status

}

var currstatus bool

func ToggleCroneHelper(code int) (error, bool) {

	var statuscheck = map[int]bool{
		0: false,
		1: true,
	}

	res, exists := statuscheck[code]

	if currstatus == res || !exists {
		return fmt.Errorf("invalid command"), false
	}

	currstatus = res

	return nil, res

}

func SalesReportHelper(code int) (time.Time, error) {

	switch code {
	case 0:
		return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, &time.Location{}), nil
	case 1:
		week := time.Now().AddDate(0, 0, -7)
		return time.Date(week.Year(), week.Month(), week.Day(), 0, 0, 0, 0, week.Location()), nil
	case 2:
		month := time.Now().AddDate(0, -1, 0)
		return time.Date(month.Year(), month.Month(), month.Day(), 0, 0, 0, 0, month.Location()), nil
	case 3:
		year := time.Now().AddDate(-1, 0, 0)
		return time.Date(year.Year(), year.Month(), year.Day(), 0, 0, 0, 0, year.Location()), nil
	default:
		return time.Time{}, fmt.Errorf("the provided timeframe doesnt exist")
	}

}
