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
	lobby.NewLobby,
	CreateHandlersMap,
)

func CreateHandlersMap(
	lobby *lobby.Lobby,
) server.InitServicesMap {
	return func() map[string]server.Handler {
		funcMap := make(map[string]server.Handler)
		funcMap["lobby"] = lobby
		return funcMap
	}
}
