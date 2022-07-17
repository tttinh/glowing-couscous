package middleware

import (
	"context"
	"encoding/json"
	"log"
	"regexp"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/tttinh/glowing-couscous/common"
	"github.com/tttinh/glowing-couscous/common/middleware/interfaces"
)

type AuthInterceptor struct {
	auth           interfaces.IJWTAuth
	ignoredMethods []string
}

func NewAuthInterceptor(auth interfaces.IJWTAuth, ignoredMethods []string) *AuthInterceptor {
	return &AuthInterceptor{
		auth:           auth,
		ignoredMethods: ignoredMethods,
	}
}

func (ai *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		method := ai.extractMethod(info.FullMethod)
		for _, m := range ai.ignoredMethods {
			if method == m {
				return handler(ctx, req)
			}
		}

		ctx, userID, err := ai.authorize(ctx)
		if err != nil {
			return nil, status.New(codes.Internal, err.Error()).Err()
		}

		// attach "x-user-id" to context
		ctx = context.WithValue(ctx, "x-user-id", userID)

		return handler(ctx, req)
	}
}

func (ai *AuthInterceptor) authorize(ctx context.Context) (context.Context, string, error) {
	m, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(m["token"]) == 0 {
		return ctx, "", status.New(codes.Unauthenticated, "missing token").Err()
	}

	authData, err := ai.auth.ValidateToken(m["token"][0])
	if err != nil {
		return ctx, "", status.New(codes.Unauthenticated, "unauthorized").Err()
	}

	var meta map[string]interface{}
	b, err := json.Marshal(authData)
	if err != nil {
		return ctx, "", status.New(codes.Unauthenticated, "unauthorized").Err()
	} else {
		if err := json.Unmarshal(b, &meta); err != nil {
			log.Println("Error while unmarshaling authData data", err)
		}
	}
	ctx = context.WithValue(ctx, "token", m["token"][0])
	ctx = common.SetContextMetadata(ctx, meta)

	return ctx, authData["sub"].(string), nil
}

func (ai *AuthInterceptor) extractMethod(fullMethod string) string {
	re := regexp.MustCompile(`.+/(\w+)$`)
	method := re.ReplaceAllString(fullMethod, "$1")

	return method
}
