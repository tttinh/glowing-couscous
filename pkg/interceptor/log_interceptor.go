package interceptor

import (
	"context"

	"google.golang.org/grpc"

	"github.com/tttinh/glowing-couscous/common/logger"
)

type LogInterceptor struct {
}

func NewLogInterceptor() *LogInterceptor {
	return &LogInterceptor{}
}

func (i *LogInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (res interface{}, err error) {
		defer func() {
			if err == nil {
				return
			}

			logger.Errorf("[%s] %v", info.FullMethod, err)
		}()

		userId, ok := ctx.Value("x-user-id").(string)
		if ok {
			logger.Infof("[%s] CALLER: %s REQ: %v", info.FullMethod, userId, req)
		} else {
			logger.Infof("[%s] REQ: %v", info.FullMethod, req)
		}

		return handler(ctx, req)
	}
}
