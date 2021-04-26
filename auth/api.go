package auth

// AuthService ...
type AuthService interface {
	Login(params *LoginParams) (customerInfo *CustomerInfo, err error)
	ReAuth(params *ReAuthParams) (err error)
}
