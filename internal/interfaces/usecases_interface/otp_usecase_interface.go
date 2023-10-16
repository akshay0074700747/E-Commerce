package usecasesinterface

type OtpUseCaseInteface interface {
	GenerateOtp(email string)(string,error)
    VerifyOtp(email, otp string) (bool)
}