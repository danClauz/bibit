package server

import (
	"context"
	"testing"

	"github.com/danClauz/bibit/bmovie/search/interfaces"
	"github.com/danClauz/bibit/bmovie/search/shared"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestServer_HealthCheck(t *testing.T) {
	assert := assert.New(t)
	logger, _ := test.NewNullLogger()

	t.Run("health check test", func(t *testing.T) {
		s := New(shared.Holder{Logger: logger}, interfaces.Holder{})

		resp, err := s.HealthCheck(context.Background(), nil)
		assert.Equal(resp.Status, "success")
		assert.NoError(err)
	})
}
