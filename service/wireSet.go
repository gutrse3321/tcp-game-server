package service

import (
	"demo/pkg/server"
	"demo/service/lobby"
	"github.com/google/wire"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/22 16:04
 * @Title:
 * --- --- ---
 * @Desc:
 */
var ProviderSet = wire.NewSet(
	lobby.ProviderSet,
	CreateHandlersMap,
)

func CreateHandlersMap(
	getLobby *lobby.GetLobby,
) server.InitServicesMap {
	return func() map[string]server.Handler {
		funcMap := make(map[string]server.Handler)
		funcMap["getLobby"] = getLobby
		return funcMap
	}
}
