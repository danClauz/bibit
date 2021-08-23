package application

import (
	"context"
	"testing"

	"github.com/danClauz/bibit/bmovie/search/external/omdb"
	omdb_mock "github.com/danClauz/bibit/bmovie/search/external/omdb/mocks"
	"github.com/danClauz/bibit/bmovie/search/repository"
	repository_mock "github.com/danClauz/bibit/bmovie/search/repository/mocks"
	"github.com/danClauz/bibit/bmovie/search/shared"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func Test_service_ListMovieByTitle(t *testing.T) {
	assert := assert.New(t)
	logger, _ := test.NewNullLogger()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOmdb := omdb_mock.NewMockClient(ctrl)
	mockRepository := repository_mock.NewMockRepository(ctrl)

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
				mockOmdb.EXPECT().SearchMovie(gomock.Any(), "request-id", "batman", 1).
					Return(&omdb.SearchResponse{}, nil)
				mockRepository.EXPECT().Store(gomock.Any(), "request-id", gomock.Any()).
					Return(nil)
			},
		},
		{
			name: "failed - external call response error",
			args: args{
				ctx:       context.Background(),
				reqId:     "request-id",
				searchKey: "batman",
				page:      1,
			},
			wantErr: true,
			mockFunc: func() {
				mockOmdb.EXPECT().SearchMovie(gomock.Any(), "request-id", "batman", 1).
					Return(nil, errors.New("an error returned"))
			},
		},
		{
			name: "failed - store error",
			args: args{
				ctx:       context.Background(),
				reqId:     "request-id",
				searchKey: "batman",
				page:      1,
			},
			wantErr: true,
			mockFunc: func() {
				mockOmdb.EXPECT().SearchMovie(gomock.Any(), "request-id", "batman", 1).
					Return(&omdb.SearchResponse{}, nil)
				mockRepository.EXPECT().Store(gomock.Any(), "request-id", gomock.Any()).
					Return(errors.New("an error returned"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()

			svc := New(repository.Holder{
				Repository: mockRepository,
			}, shared.Holder{
				Logger: logger,
			}, mockOmdb)

			_, err := svc.ListMovieByTitle(tt.args.ctx, tt.args.reqId, tt.args.searchKey, tt.args.page)
			assert.Equal(tt.wantErr, err != nil)
		})
	}
}

func Test_service_MovieDetailByImdbId(t *testing.T) {
	assert := assert.New(t)
	logger, _ := test.NewNullLogger()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOmdb := omdb_mock.NewMockClient(ctrl)
	mockRepository := repository_mock.NewMockRepository(ctrl)

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
				searchKey: "omdb-id",
			},
			wantErr: false,
			mockFunc: func() {
				mockOmdb.EXPECT().SearchMovieByImdbId(gomock.Any(), "request-id", "omdb-id").
					Return(&omdb.SearchByTitleResponse{}, nil)
				mockRepository.EXPECT().Store(gomock.Any(), "request-id", gomock.Any()).
					Return(nil)
			},
		},
		{
			name: "failed - external call response error",
			args: args{
				ctx:       context.Background(),
				reqId:     "request-id",
				searchKey: "omdb-id",
			},
			wantErr: true,
			mockFunc: func() {
				mockOmdb.EXPECT().SearchMovieByImdbId(gomock.Any(), "request-id", "omdb-id").
					Return(nil, errors.New("an error returned"))
			},
		},
		{
			name: "failed - store error",
			args: args{
				ctx:       context.Background(),
				reqId:     "request-id",
				searchKey: "omdb-id",
			},
			wantErr: true,
			mockFunc: func() {
				mockOmdb.EXPECT().SearchMovieByImdbId(gomock.Any(), "request-id", "omdb-id").
					Return(&omdb.SearchByTitleResponse{}, nil)
				mockRepository.EXPECT().Store(gomock.Any(), "request-id", gomock.Any()).
					Return(errors.New("an error returned"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()

			svc := New(repository.Holder{
				Repository: mockRepository,
			}, shared.Holder{
				Logger: logger,
			}, mockOmdb)

			_, err := svc.MovieDetailByImdbId(tt.args.ctx, tt.args.reqId, tt.args.searchKey)
			assert.Equal(tt.wantErr, err != nil)
		})
	}
}
