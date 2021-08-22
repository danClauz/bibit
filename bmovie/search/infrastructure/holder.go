package infrastructure

import (
	"fmt"
	searchpb "github.com/danClauz/bibit/bmovie/search/gen"
	"github.com/danClauz/bibit/bmovie/search/infrastructure/gateway"
	"github.com/danClauz/bibit/bmovie/search/infrastructure/server"
	"github.com/danClauz/bibit/bmovie/search/shared"
	"go.uber.org/dig"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"time"
)

type Holder struct {
	dig.In
	Gateway *gateway.Gateway
	Server  *server.Server
	Sh      shared.Holder
}

func (h *Holder) ServeHttp() {
	RegisterDefaultMiddleware(h)
	shared := h.Sh
	logger := shared.Logger

	bmovie := shared.Echo.Group("/bmovie")

	//bmovie.GET("/health-check", h.Controllers.HealthCheck)

	v1 := bmovie.Group("/v1")

	home := v1.Group("/")
	home.GET("", h.Gateway.SearchMovie)

	serverCfg := shared.Config.HttpServer

	s := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", serverCfg.Host, "8080"),
		ReadTimeout:  serverCfg.ReadTimeout * time.Second,
		WriteTimeout: serverCfg.WriteTimeout * time.Second,
	}

	ch := make(chan error, 1)
	go func() {
		ch <- shared.Echo.StartServer(s)
	}()

	logger.Info("Failed to serve:", <-ch)
	close(ch)

	logger.Info("Server interrupted!")
}

func (h *Holder) ServeGateway() {
	logger := h.Sh.Logger

	ch := make(chan error, 1)

	go func() {
		ch <- h.Gateway.RunGateway()
	}()

	logger.Info("Failed to serve:", <-ch)
	close(ch)

	logger.Info("Server interrupted!")
}

func (h *Holder) ServeGrpc() {
	logger := h.Sh.Logger
	grpcServer := h.Sh.Config.GrpcServer

	address := fmt.Sprintf("%s:%s", grpcServer.Host, grpcServer.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		logger.Fatalf("error to listen to %s. err: %v", address, err)
	}

	logger.Infof("gRPC server is listening on %v", address)

	s := grpc.NewServer()
	searchpb.RegisterSearchServiceServer(s, h.Server)

	ch := make(chan error, 1)

	go func() {
		ch <- s.Serve(lis)
	}()

	logger.Info("Failed to serve:", <-ch)
	close(ch)

	logger.Info("Server interrupted!")

	logger.Info("Stopping the server")
	s.Stop()

	logger.Info("Closing the listener")
	_ = lis.Close()
}
