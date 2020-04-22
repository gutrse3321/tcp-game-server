//+build wireinject

package main

import (
	"demo/pkg/config"
	"demo/pkg/server"
	"demo/service"
	"github.com/google/wire"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/22 15:44
 * @Title:
 * --- --- ---
 * @Desc:
 */
var providerSet = wire.NewSet(
	config.ProviderSet,
	server.ProviderSet,
	//database.ProviderSet,
	//redis.ProviderSet,
	service.ProviderSet,
)

func CreateServer(configPath string) (*server.Server, error) {
	panic(wire.Build(providerSet))
}
