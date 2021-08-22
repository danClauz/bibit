package gateway

import (
	"context"
	"fmt"
	searchpb "github.com/danClauz/bibit/bmovie/search/gen"
	"github.com/danClauz/bibit/bmovie/search/interfaces"
	"github.com/danClauz/bibit/bmovie/search/shared"
	"github.com/danClauz/bibit/bmovie/search/shared/logtag"
	"github.com/danClauz/bibit/bmovie/search/shared/response"
	"github.com/danClauz/bibit/bmovie/search/shared/utils"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"strconv"
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

func (g *Gateway) RunGateway() error {
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

	grpcServer := g.sh.Config.GrpcServer
	httpServer := g.sh.Config.HttpServer

	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := searchpb.RegisterSearchServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%s:%s", grpcServer.Host, grpcServer.Port), opts); err != nil {
		return err
	}

	return http.ListenAndServe(fmt.Sprintf("%s:%s", httpServer.Host, httpServer.Port), mux)
}

func (g *Gateway) SearchMovie(ctx echo.Context) error {
	reqId := utils.FetchRequestIdFromContext(ctx)

	search := ctx.QueryParam("s")
	page := ctx.QueryParam("p")
	id := ctx.QueryParam("i")

	logger := g.sh.Logger.WithFields(map[string]interface{}{})

	logger.Infof(logtag.RequestTmpl, logtag.PerformSearchMovie)

	var resp interface{}
	var respErr error = nil

	if search != "" {
		p, err := strconv.Atoi(page)
		if err != nil {
			p = 1
		}

		resp, respErr = g.ih.SearchMovie(ctx.Request().Context(), reqId, search, p)
	} else if id != "" {
		resp, respErr = g.ih.DetailMovie(ctx.Request().Context(), reqId, id)
	} else {
		err := response.Error(response.BAD_REQUEST, errors.New("invalid query input"))
		return response.Body(ctx, nil, err)
	}

	if respErr != nil {
		logger.Errorf(logtag.ErrorTmpl, logtag.PerformSearchMovie, respErr)
		return response.Body(ctx, nil, respErr)
	}

	return response.Body(ctx, resp, nil)
}
