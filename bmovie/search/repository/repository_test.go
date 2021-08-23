package repository

import (
	"context"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/danClauz/bibit/bmovie/search/model"
	"github.com/danClauz/bibit/bmovie/search/shared"

	"testing"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func Test_repo_Store(t *testing.T) {
	assert := assert.New(t)
	logger, _ := test.NewNullLogger()

	type args struct {
		ctx   context.Context
		reqId string
		obj   interface{}
	}
	tests := []struct {
		name    string
		args    args
		mockErr error
		wantErr bool
	}{
		{
			name: "success - search logs",
			args: args{
				ctx:   context.Background(),
				reqId: "request-id",
				obj:   &model.SearchHistory{},
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "success - detail log",
			args: args{
				ctx:   context.Background(),
				reqId: "request-id",
				obj:   &model.SearchHistory{},
			},
			mockErr: nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, _, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			//switch tt.args.obj.(type) {
			//case model.SearchLog:
			//	mock.ExpectExec("INSERT INTO search_logs").WillReturnResult(sqlmock.NewResult(1, 1))
			//case model.DetailLog:
			//	mock.ExpectExec("INSERT INTO detail_logs").WillReturnResult(sqlmock.NewResult(1, 1))
			//}

			//gormdb, err := gorm.Open(mysql.New(mysql.Config{
			//	Conn: db,
			//}), &gorm.Config{})
			//assert.NoError(err)

			r := New(shared.Holder{
				Logger: logger,
				//Mysql:  gormdb,
			})

			assert.NotNil(r)

			err = r.Store(tt.args.ctx, tt.args.reqId, tt.args.obj)
			assert.Equal(tt.wantErr, err != nil)
		})
	}
}
