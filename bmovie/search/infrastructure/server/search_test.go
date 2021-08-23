package server

import (
	"context"
	"testing"

	"github.com/danClauz/bibit/bmovie/search/external/omdb"
	searchpb "github.com/danClauz/bibit/bmovie/search/gen"
	"github.com/danClauz/bibit/bmovie/search/interfaces"
	interfaces_mock "github.com/danClauz/bibit/bmovie/search/interfaces/mocks"
	"github.com/danClauz/bibit/bmovie/search/model"
	"github.com/danClauz/bibit/bmovie/search/shared"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestServer_DetailMovie(t *testing.T) {
	assert := assert.New(t)
	logger, _ := test.NewNullLogger()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockInterfaces := interfaces_mock.NewMockInterfaces(ctrl)

	type args struct {
		ctx context.Context
		req *searchpb.DetailMovieRequest
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
				ctx: context.Background(),
				req: &searchpb.DetailMovieRequest{
					ImdbId: "imdb-id",
				},
			},
			wantErr: false,
			mockFunc: func() {
				mockInterfaces.EXPECT().DetailMovie(gomock.Any(), gomock.Any(), "imdb-id").
					Return(&model.SearchHistory{
						Result: &omdb.SearchByTitleResponse{},
					}, nil)
			},
		},
		{
			name: "failed",
			args: args{
				ctx: context.Background(),
				req: &searchpb.DetailMovieRequest{
					ImdbId: "imdb-id",
				},
			},
			wantErr: true,
			mockFunc: func() {
				mockInterfaces.EXPECT().DetailMovie(gomock.Any(), gomock.Any(), "imdb-id").
					Return(nil, errors.New("something went wrong"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()

			s := New(shared.Holder{
				Logger: logger,
			}, interfaces.Holder{
				Interfaces: mockInterfaces,
			})

			_, err := s.DetailMovie(tt.args.ctx, tt.args.req)
			assert.Equal(tt.wantErr, err != nil)
		})
	}
}

func TestServer_SearchMovie(t *testing.T) {
	assert := assert.New(t)
	logger, _ := test.NewNullLogger()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockInterfaces := interfaces_mock.NewMockInterfaces(ctrl)

	type args struct {
		ctx context.Context
		req *searchpb.SearchMovieRequest
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
				ctx: context.Background(),
				req: &searchpb.SearchMovieRequest{
					Search: "batman",
					Page:   1,
				},
			},
			wantErr: false,
			mockFunc: func() {
				mockInterfaces.EXPECT().SearchMovie(gomock.Any(), gomock.Any(), "batman", 1).
					Return(&model.SearchHistory{
						Result: &omdb.SearchByTitleResponse{},
					}, nil)
			},
		},
		{
			name: "failed",
			args: args{
				ctx: context.Background(),
				req: &searchpb.SearchMovieRequest{
					Search: "batman",
					Page:   1,
				},
			},
			wantErr: true,
			mockFunc: func() {
				mockInterfaces.EXPECT().SearchMovie(gomock.Any(), gomock.Any(), "batman", 1).
					Return(nil, errors.New("something went wrong"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()

			s := New(shared.Holder{
				Logger: logger,
			}, interfaces.Holder{
				Interfaces: mockInterfaces,
			})

			_, err := s.SearchMovie(tt.args.ctx, tt.args.req)
			assert.Equal(tt.wantErr, err != nil)
		})
	}
}
