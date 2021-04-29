package auth

import (
	"http/web"
)

const (
	errUserNotFound  = web.Unauthorized("Tài khoản không tồn tại")
	errUserNotVerify = web.Unauthorized("Tài khoản chưa được xác thực qua email")
)
