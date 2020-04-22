package database

import (
	"fmt"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/21 18:25
 * @Title:
 * --- --- ---
 * @Desc:
 */
type Options struct {
	User         string
	Password     string
	Host         string
	Db           string
	MaxIdleConns int
	MaxOpenConns int
	SearchUrl    string
	Debug        bool
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error

	opt := &Options{}
	if err = v.UnmarshalKey("database", opt); err != nil {
		return nil, errors.Wrap(err, "Unmarshal database config error")
	}

	log.Info("load database config success")

	return opt, err
}

/**
初始化连接数据库(mySql)
*/
func New(opt *Options) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", opt.User, opt.Password, opt.Host, opt.Db, opt.SearchUrl))
	if err != nil {
		return nil, errors.Wrap(err, "gorm open database connection error")
	}
	if opt.Debug {
		db = db.Debug()
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(opt.MaxIdleConns)
	db.DB().SetMaxOpenConns(opt.MaxOpenConns)

	return db, nil
}

var ProviderSet = wire.NewSet(New, NewOptions)
