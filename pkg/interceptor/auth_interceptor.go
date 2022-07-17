package interceptor

import (
	"github.com/tttinh/glowing-couscous/common/middleware"
	"github.com/tttinh/glowing-couscous/common/middleware/identity"
)

var publicMethods = []string{}

func NewAuthInterceptor(authKey string) *middleware.AuthInterceptor {
	auth := identity.NewAuth(authKey)
	return middleware.NewAuthInterceptor(auth, publicMethods)
}
