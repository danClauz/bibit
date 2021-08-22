package di

import (
	"github.com/danClauz/bibit/bmovie/search/shared/config"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(conf *config.EnvConfig) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(conf.Mysql.Master.DSN), &gorm.Config{})

	//if err := db.AutoMigrate(&model.SearchHistory{}); err != nil {
	//	return nil, err
	//}

	return db, err
}

func init() {
	if err := Container.Provide(NewMysql); err != nil {
		panic(errors.Wrap(err, "failed to provide mysql"))
	}
}
