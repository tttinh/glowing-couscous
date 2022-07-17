package main

import (
	"fmt"
	"net"

	"github.com/getsentry/sentry-go"
	"github.com/tttinh/glowing-couscous/common/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	genpb "github.com/tttinh/glowing-couscous/api/genproto/go"

	"github.com/tttinh/glowing-couscous/internal/company"
	"github.com/tttinh/glowing-couscous/internal/company/adapter/client"
	"github.com/tttinh/glowing-couscous/pkg/config"
	"github.com/tttinh/glowing-couscous/pkg/db"
	"github.com/tttinh/glowing-couscous/pkg/interceptor"
)

func main() {
	conf := config.NewConfig()
	logger.Initialize(conf.Env)

	err := initSentry(conf)
	if err != nil {
		logger.Fatal("sentry.Init: ", err)

	}

	db, err := db.Init(conf.Env, conf.MongoURL, conf.MongoDatabase)
	if err != nil {
		logger.Fatal("db.Init: ", err)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.ServicePort))
	if err != nil {
		logger.Fatal("listener.Init: ", err)
	}

	authInterceptor := interceptor.NewAuthInterceptor(conf.SecretAuthKey)
	logInterceptor := interceptor.NewLogInterceptor()
	recoveryInterceptor := interceptor.NewRecoveryInterceptor()

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(authInterceptor.Unary()),
		grpc.ChainUnaryInterceptor(logInterceptor.Unary()),
		grpc.ChainUnaryInterceptor(recoveryInterceptor.Unary()),
	)

	userClient := client.NewAuth(conf.AuthServiceEndpoint)
	companyServer := company.New(userClient, db)
	genpb.RegisterCompanyServiceServer(grpcServer, companyServer)
	reflection.Register(grpcServer)

	logger.Info("Server started ...")

	grpcServer.Serve(listener)
}

func initSentry(conf *config.Config) error {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:        conf.SentryDSN,
		ServerName: "banking-service",
	})

	if err != nil {
		return err
	}

	return nil
}
