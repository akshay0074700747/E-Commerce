package middlewares

import (
	"fmt"
	"testing"
)

type OtpMock struct {
	email string
	error error
}

var TestCases = []OtpMock{
	{email: "ask029849@gmail.com", error: nil},
	{email: "sidx141", error: fmt.Errorf("this is not a valid email")},
	{email: "sidx141202@gmail.com", error: nil},
	{email: "sidx141202@gmacom", error: fmt.Errorf("this is not a valid email")},
	{email: "sharoonkp267@gmail.com", error: nil},
}

func TestOtp_Gen(t *testing.T) {

	for _, test := range TestCases {
		if err := Otp_Gen(test.email); err != test.error {
			t.Fail()
		}
	}

}

type OtpVerifyMock struct {
	email   string
	otp     string
	error   error
	success bool
}

var VerifyTestcases = []OtpVerifyMock{
	{email: "ask029849@gmail.com", error: nil, otp: "438381", success: true},
	{email: "sidx141202@gmail.com", error: nil, otp: "1465", success: false},
	{email: "sharoonkp267@gmail.com", error: nil, otp: "963792", success: true},
}

func TestOtp_Verify(t *testing.T) {

	for _, test := range VerifyTestcases {
		if err, res := Otp_Verify(test.email, test.otp); err != test.error || res != test.success {
			t.Fail()
		}
	}

}
