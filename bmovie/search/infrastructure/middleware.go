package infrastructure

import (
	"fmt"
	"github.com/danClauz/bibit/bmovie/search/shared/utils"
	"github.com/labstack/echo/v4/middleware"
)

var (
	mdLogFmt = fmt.Sprintf("time=\"${time_rfc3339}\" level=\"echo\" x-request-id=\"${header:%s}\" remote_ip=\"${remote_ip}\" host=\"${host}\" method=\"${method}\" uri=\"${uri}\" user_agent=\"${user_agent}\" status=\"${status}\" latency_human=\"${latency_human}\" bytes_in=\"${bytes_in}\" bytes_out=\"${bytes_out}\"\n", utils.XRequestId)
)

func RegisterDefaultMiddleware(holder *Holder) {
	// - register your default middleware here
	holder.Sh.Echo.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: mdLogFmt,
		}),
	)
	holder.Sh.Echo.Use(middleware.CORS())
}
