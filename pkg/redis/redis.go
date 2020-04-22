package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
	"time"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/21 18:27
 * @Title:
 * --- --- ---
 * @Desc:
 */
type Options struct {
	Host        string
	Port        int
	Password    string
	Db          int
	Timeout     time.Duration
	ExpiredTime int
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error

	opt := &Options{}
	if err = v.UnmarshalKey("redis", opt); err != nil {
		return nil, errors.Wrap(err, "Unmarshal redis config error")
	}

	log.Info("load redis config success")

	return opt, err
}

/**
初始化连接数据库(redis)
*/
func New(opt *Options) (*redis.Client, error) {
	redisOpt := &redis.Options{
		Addr:         fmt.Sprintf("%s:%d", opt.Host, opt.Port),
		Password:     opt.Password,
		DB:           opt.Db,
		ReadTimeout:  opt.Timeout * time.Second,
		WriteTimeout: opt.Timeout * time.Second,
	}

	client := redis.NewClient(redisOpt)
	if _, err := client.Ping().Result(); err != nil {
		return nil, errors.Wrap(err, "redis create client error")
	}

	return client, nil
}

var ProviderSet = wire.NewSet(New, NewOptions)
