package lobby

import (
	"github.com/panjf2000/gnet"
	"log"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/22 15:55
 * @Title:
 * --- --- ---
 * @Desc:
 */
type GetLobby struct {
	//db    *gorm.DB
	//redis *redis.Client
}

//func NewGetLobby(db *gorm.DB, redis *redis.Client) *GetLobby {
//	return &GetLobby{db, redis}
//}

func NewGetLobby() *GetLobby {
	return &GetLobby{}
}

func (h *GetLobby) Handle(ctx gnet.Conn, json string) error {
	log.Println("in handler")
	return nil
}
