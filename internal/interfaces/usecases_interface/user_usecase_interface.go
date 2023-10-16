package usecasesinterface

type UserUsecaseInterface interface {
	RegisterUser (email,password,mobile,name string) error
}