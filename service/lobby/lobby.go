package lobby

import (
	"encoding/json"
	"github.com/panjf2000/gnet"
	"log"
	"reflect"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/22 15:55
 * @Title:
 * --- --- ---
 * @Desc:
 */
type Lobby struct {
	//db    *gorm.DB
	//redis *redis.Client
}

//func NewGetLobby(db *gorm.DB, redis *redis.Client) *GetLobby {
//	return &GetLobby{db, redis}
//}

func NewLobby() *Lobby {
	return &Lobby{}
}

type Model struct {
	NickName string
}

func (h *Lobby) Handle(ctx gnet.Conn, service string, jsonStr []byte) error {
	params := make([]reflect.Value, 2)
	params[0] = reflect.ValueOf(ctx)
	params[1] = reflect.ValueOf(jsonStr)

	valueOf := reflect.ValueOf(h)
	valueOf.MethodByName(service).Call(params)
	return nil
}

func (h *Lobby) GetLobby(ctx gnet.Conn, body []byte) {
	model := &Model{}
	json.Unmarshal(body, model)
	log.Println("in handler:", model.NickName)
	ctx.AsyncWrite([]byte("handler call:" + model.NickName))
}
