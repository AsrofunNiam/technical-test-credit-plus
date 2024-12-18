package service

import "github.com/AsrofunNiam/technical-test-credit-plus/model/web"

type UserService interface {
	Login(identity, password, userAgent, remoteAddress *string, request *web.UserLoginRequest) web.TokenResponse
}
