package main

import (
	"context"
	stdLog "log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aldisaputra17/go-micro/src/domain/pb"
	"github.com/aldisaputra17/go-micro/src/gapi"
	"github.com/aldisaputra17/go-micro/src/module"
	"github.com/aldisaputra17/go-micro/toolkit/config"
	"github.com/aldisaputra17/go-micro/toolkit/db"
	"github.com/aldisaputra17/go-micro/toolkit/echokit"
	"github.com/aldisaputra17/go-micro/toolkit/log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	var err error

	setDefaultTimezone()

	ctx, cancel := config.NewRuntimeContext()
	defer func() {
		cancel()

		if err != nil {
			log.FromCtx(ctx).Error(err, "found error")
		}
	}()

	if os.Getenv("APP_ENV") == "" {
		if err = godotenv.Load(); err != nil {
			stdLog.Fatalf("error load env file : %s", err.Error())
		}
	}

	// setup logger
	logger, err := log.NewFromConfig()
	if err != nil {
		stdLog.Fatalf("error setup log : %s", err.Error())
		return
	}

	logger.Set()

	database, err := db.NewDatabaseConnection()
	if err != nil {
		return
	}

	mdl := module.NewModule(database)

	go runGRPCGateway(os.Getenv("GRPC_ADDRESS"), mdl)
	runGRPCServer(os.Getenv("GRPC_ADDRESS"), mdl)
}

func runAPIServer(ctx context.Context, mdl *module.Module) {
	echokit.RunEchoHTTP(ctx, mdl)
}

func runGRPCServer(address string, mdl *module.Module) {
	ctx := context.Background()

	server, err := gapi.NewGRPCServer(address, mdl)
	if err != nil {
		log.FromCtx(ctx).Error(err, "cannot create server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", server.Address)
	if err != nil {
		log.FromCtx(ctx).Error(err, "cannot create listener:", err)
	}

	log.Printf("start gRPC server on %s", listener.Addr().String())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Printf("Stopping gRPC server on port %s", listener.Addr().String())
		grpcServer.GracefulStop()
	}()

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("cannot start gRPC server")
	}
}

func runGRPCGateway(address string, mdl *module.Module) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server, err := gapi.NewGRPCServer(address, mdl)
	if err != nil {
		log.FromCtx(ctx).Error(err, "cannot create server:", err)
	}

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOption)
	if err := pb.RegisterUserServiceHandlerServer(ctx, grpcMux, server); err != nil {
		log.Fatal("cannot register server handler:", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", os.Getenv("HTTP_ADDRESS"))
	if err != nil {
		log.FromCtx(ctx).Error(err, "cannot create listener")
	}

	log.Printf("start HTTP server on %s", listener.Addr().String())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Printf("Stopping HTTP server on port %s", listener.Addr().String())
	}()

	if err := http.Serve(listener, mux); err != nil {
		log.Fatal("cannot start HTTP Gateway server:", err)
	}
}

func setDefaultTimezone() {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		loc = time.Now().Location()
	}

	time.Local = loc
}
