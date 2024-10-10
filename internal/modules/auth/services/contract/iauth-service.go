package contract

type IAuthService interface {
	Registration(string, string) error
	Login() error
	Logout(string) error
}
