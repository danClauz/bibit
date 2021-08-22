package interfaces

import (
	"context"
	"github.com/danClauz/bibit/bmovie/search/application"
	application_mock "github.com/danClauz/bibit/bmovie/search/application/mocks"
	"github.com/danClauz/bibit/bmovie/search/model"
	"github.com/danClauz/bibit/bmovie/search/shared"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_service_DetailMovie(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAppHolder := application_mock.NewMockApplication(ctrl)

	type args struct {
		ctx       context.Context
		reqId     string
		searchKey string
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		mockFunc func()
	}{
		{
			name: "success",
			args: args{
				ctx:       context.Background(),
				reqId:     "request-id",
				searchKey: "batman",
			},
			wantErr: false,
			mockFunc: func() {
				mockAppHolder.EXPECT().MovieDetailByImdbId(gomock.Any(), "request-id", "batman").
					Return(&model.SearchHistory{}, nil)
			},
		},
		{
			name: "failed - app layer return error",
			args: args{
				ctx:       context.Background(),
				reqId:     "request-id",
				searchKey: "batman",
			},
			wantErr: true,
			mockFunc: func() {
				mockAppHolder.EXPECT().MovieDetailByImdbId(gomock.Any(), "request-id", "batman").
					Return(nil, errors.New("something went wrong"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()

			impl := New(shared.Holder{}, application.Holder{
				Application: mockAppHolder,
			})

			assert.NotNil(impl)

			_, err := impl.DetailMovie(tt.args.ctx, tt.args.reqId, tt.args.searchKey)
			assert.Equal(tt.wantErr, err != nil)
		})
	}
}

func Test_service_SearchMovie(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAppHolder := application_mock.NewMockApplication(ctrl)

	type args struct {
		ctx       context.Context
		reqId     string
		searchKey string
		page      int
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		mockFunc func()
	}{
		{
			name: "success",
			args: args{
				ctx:       context.Background(),
				reqId:     "request-id",
				searchKey: "batman",
				page:      1,
			},
			wantErr: false,
			mockFunc: func() {
				mockAppHolder.EXPECT().ListMovieByTitle(gomock.Any(), "request-id", "batman", 1).
					Return(&model.SearchHistory{}, nil)
			},
		},
		{
			name: "failed - app layer return error",
			args: args{
				ctx:       context.Background(),
				reqId:     "request-id",
				searchKey: "batman",
				page:      1,
			},
			wantErr: true,
			mockFunc: func() {
				mockAppHolder.EXPECT().ListMovieByTitle(gomock.Any(), "request-id", "batman", 1).
					Return(nil, errors.New("something went wrong"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()

			impl := New(shared.Holder{}, application.Holder{
				Application: mockAppHolder,
			})

			assert.NotNil(impl)

			_, err := impl.SearchMovie(tt.args.ctx, tt.args.reqId, tt.args.searchKey, tt.args.page)
			assert.Equal(tt.wantErr, err != nil)
		})
	}
}
