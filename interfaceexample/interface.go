package interfaceExample

type IUserService interface {
	Register()
	Login()
	Logout()
	Edit()
}

// act like implement IUserService
type MemoryUserService struct {
	//...
	//Users []User
}

// act like implement IUserService
type DatabaseUserService struct {
	// ...
	//Mongodb *SomeDatabaseConnection
}

type UserController struct {
	Service *IUserService
}
