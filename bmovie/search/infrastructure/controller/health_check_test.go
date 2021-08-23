package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danClauz/bibit/bmovie/search/interfaces"
	"github.com/danClauz/bibit/bmovie/search/shared"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestController_HealthCheck(t *testing.T) {
	assert := assert.New(t)
	logger, _ := test.NewNullLogger()

	t.Run("health check test", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		ctrl := New(shared.Holder{Logger: logger}, interfaces.Holder{})
		assert.NotNil(ctrl)

		if assert.NoError(ctrl.HealthCheck(ctx)) {
			assert.Equal(http.StatusOK, rec.Code)
		}
	})
}
