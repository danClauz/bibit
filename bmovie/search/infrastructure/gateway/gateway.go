package gateway

import (
	"context"
	"fmt"
	"net/http"

	searchpb "github.com/danClauz/bibit/bmovie/search/gen"
	"github.com/danClauz/bibit/bmovie/search/interfaces"
	"github.com/danClauz/bibit/bmovie/search/shared"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

type Gateway struct {
	sh shared.Holder
	ih interfaces.Holder
}

func New(sh shared.Holder, ih interfaces.Holder) *Gateway {
	return &Gateway{
		sh: sh,
		ih: ih,
	}
}

func (gtw *Gateway) RunGateway() error {
	grpcServer := gtw.sh.Config.GrpcServer
	grpcGateway := gtw.sh.Config.GrpcGateway

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:   true,
					EmitUnpopulated: true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		}),
	)

	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := searchpb.RegisterSearchServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%s:%s", grpcServer.Host, grpcServer.Port), opts); err != nil {
		return err
	}

	return http.ListenAndServe(fmt.Sprintf("%s:%s", grpcGateway.Host, grpcGateway.Port), mux)
}
