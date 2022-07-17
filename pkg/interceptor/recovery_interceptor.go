package interceptor

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/tttinh/glowing-couscous/common/logger"
	"google.golang.org/grpc"

	ce "github.com/tttinh/glowing-couscous/common/errors"
)

type RecoveryInterceptor struct {
}

func NewRecoveryInterceptor() *RecoveryInterceptor {
	return &RecoveryInterceptor{}
}

func (i *RecoveryInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (res interface{}, err error) {
		defer func() {
			if pErr := recover(); pErr != nil {
				sentry.ConfigureScope(func(scope *sentry.Scope) {
					scope.SetExtra("context", "panic recover")
					scope.SetExtra("request function", info.FullMethod)
					scope.SetExtra("request", req)
					sentry.CaptureException(pErr.(error))
				})
				logger.Info("Recovered from err: ", pErr)
				err = ce.New(ce.ErrSomethingWentWrong)
				res = nil
			}
		}()

		return handler(ctx, req)
	}
}
