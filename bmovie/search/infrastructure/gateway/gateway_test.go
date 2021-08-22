package gateway

import (
	"fmt"
	"github.com/danClauz/bibit/bmovie/search/interfaces"
	interfaces_mock "github.com/danClauz/bibit/bmovie/search/interfaces/mocks"
	"github.com/danClauz/bibit/bmovie/search/model"
	"github.com/danClauz/bibit/bmovie/search/shared"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestControllers_SearchMovie(t *testing.T) {
	assert := assert.New(t)
	logger, _ := test.NewNullLogger()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockInterface := interfaces_mock.NewMockInterfaces(ctrl)

	type args struct {
		search string
		page   interface{}
		id     string
	}

	tests := []struct {
		name           string
		args           args
		expectedStatus int
		mockFunc       func()
	}{
		{
			name: "success - search 200",
			args: args{
				search: "batman",
				page:   1,
				id:     "",
			},
			expectedStatus: http.StatusOK,
			mockFunc: func() {
				mockInterface.EXPECT().SearchMovie(gomock.Any(), gomock.Any(), "batman", 1).
					Return(&model.SearchHistory{}, nil).Times(1)
			},
		},
		{
			name: "success - search 200 - invalid page param",
			args: args{
				search: "batman",
				page:   "abcd",
				id:     "",
			},
			expectedStatus: http.StatusOK,
			mockFunc: func() {
				mockInterface.EXPECT().SearchMovie(gomock.Any(), gomock.Any(), "batman", 1).
					Return(&model.SearchHistory{}, nil).Times(1)
			},
		},
		{
			name: "success - detail 200",
			args: args{
				search: "",
				page:   0,
				id:     "imdb-id",
			},
			expectedStatus: http.StatusOK,
			mockFunc: func() {
				mockInterface.EXPECT().DetailMovie(gomock.Any(), gomock.Any(), "imdb-id").
					Return(&model.SearchHistory{}, nil).Times(1)
			},
		},
		{
			name:           "failed - 400",
			args:           args{},
			expectedStatus: http.StatusBadRequest,
			mockFunc:       func() {},
		},
		{
			name: "failed - 500",
			args: args{
				search: "",
				page:   0,
				id:     "imdb-id",
			},
			expectedStatus: http.StatusInternalServerError,
			mockFunc: func() {
				mockInterface.EXPECT().DetailMovie(gomock.Any(), gomock.Any(), "imdb-id").
					Return(nil, errors.New("something went wrong")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/?s=%s&p=%d&i=%s", tt.args.search, tt.args.page, tt.args.id), nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			c := New(shared.Holder{
				Logger: logger,
			}, interfaces.Holder{
				Interfaces: mockInterface,
			})
			assert.NotNil(c)

			if assert.NoError(c.SearchMovie(ctx)) {
				assert.Equal(tt.expectedStatus, rec.Code)
			}
		})
	}
}

