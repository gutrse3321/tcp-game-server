package lobby

import "github.com/google/wire"

/**
 * @Author: Tomonori
 * @Date: 2020/4/22 16:03
 * @Title:
 * --- --- ---
 * @Desc:
 */
var ProviderSet = wire.NewSet(
	NewGetLobby,
)
