package client

import (
	"context"
	"time"

	genpb "github.com/tttinh/glowing-couscous/api/genproto/go"

	"github.com/tttinh/glowing-couscous/internal/company/domain"

	"github.com/tttinh/glowing-couscous/common/errors"
	"github.com/tttinh/glowing-couscous/common/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type User struct {
	endpoint string
}

func NewAuth(endpoint string) *User {
	return &User{endpoint}
}

func (c *User) GetUserProfile(ctx context.Context) (*domain.Employee, error) {
	conn, err := grpc.DialContext(ctx, c.endpoint, grpc.WithInsecure())
	if err != nil {
		logger.Error("Cannot connect to service: ", err)
		return nil, err
	}
	defer conn.Close()

	client := genpb.NewUserServiceClient(conn)

	m, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(m["token"]) == 0 {
		return nil, errors.New(errors.ErrUnauthenticated, "missing token")
	}

	newCtx := metadata.NewOutgoingContext(ctx, m)
	r, err := client.GetMe(newCtx, &genpb.UserGetMeReq{})

	if err != nil {
		logger.Error("User.GetMe failed: ", err)
		return nil, err
	}

	return &domain.Employee{
		Id:        r.Id,
		Name:      r.Name,
		Email:     r.Email,
		Phone:     r.PhoneCode + r.PhoneNumber,
		Avatar:    r.Avatar,
		LastLogin: getTimeFromInt64(r.LastLogin),
	}, nil
}

func getTimeFromInt64(timestamp int64) *time.Time {
	if timestamp > 0 {
		t := time.Unix(0, timestamp*int64(time.Millisecond))
		return &t
	}

	return nil
}
