package middleware

import (
	"context"
	"log"

	"github.com/getsentry/sentry-go"
	"google.golang.org/grpc"
)

type PanicRecoveryInterceptor struct {
}

func NewPanicRecoveryInterceptor() *PanicRecoveryInterceptor {
	return &PanicRecoveryInterceptor{}
}

func (pri *PanicRecoveryInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		defer func() {
			if err := recover(); err != nil {
				sentry.ConfigureScope(func(scope *sentry.Scope) {
					scope.SetExtra("context", "panic recover")
					scope.SetExtra("request function", info.FullMethod)
					scope.SetExtra("request", req)
					sentry.CaptureException(err.(error))
				})
				log.Println("Recovered from err: ", err)
			}
		}()

		return handler(ctx, req)
	}
}
