package goboond

type AuthenticationMethod string

const (
	AuthenticationMethodClient AuthenticationMethod = "client"
	AuthenticationMethodApp    AuthenticationMethod = "app"
	AuthenticationMethodBasic  AuthenticationMethod = "basic"
)
