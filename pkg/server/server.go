package server

import (
	"encoding/json"
	"fmt"
	"github.com/google/wire"
	"github.com/panjf2000/gnet"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"log"
	"sync"
	"time"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/22 12:27
 * @Title:
 * --- --- ---
 * @Desc:
 */
type Options struct {
	Port         int
	MultiCore    bool
	BeatInterval time.Duration
}

func NewOptions(v *viper.Viper) (*Options, error) {
	opt := &Options{}

	if err := v.UnmarshalKey("server", opt); err != nil {
		return nil, err
	}

	return opt, nil
}

type TcpServer struct {
	*gnet.EventServer
	tick             time.Duration
	connectedSockets sync.Map
	port             int
	multicore        bool
	handlers         map[string]Handler
}

func (s *TcpServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	log.Printf("Echo server is listening on %s (multi-cores: %t, loops: %d)\n",
		srv.Addr.String(), srv.Multicore, srv.NumEventLoop)
	return
}

func (s *TcpServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	log.Printf("Socket with addr: %s has been opened...\n", c.RemoteAddr().String())
	c.AsyncWrite([]byte("[Server] connected"))
	s.connectedSockets.Store(c.RemoteAddr().String(), c)
	return
}
func (s *TcpServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	log.Printf("Socket with addr: %s is closing...\n", c.RemoteAddr().String())
	s.connectedSockets.Delete(c.RemoteAddr().String())
	return
}

func (s *TcpServer) Tick() (delay time.Duration, action gnet.Action) {
	log.Println("heart beat ... ")
	s.connectedSockets.Range(func(key, value interface{}) bool {
		addr := key.(string)
		c := value.(gnet.Conn)
		c.AsyncWrite([]byte(fmt.Sprintf("heart beating to %s\n", addr)))
		return true
	})
	delay = s.tick
	return
}

type Handler interface {
	Handle(ctx gnet.Conn, service string, jsonStr []byte) error
}

type ClientArgs struct {
	Handler string
	Service string
	Model   string
}

func (s *TcpServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	data := append([]byte{}, frame...)
	go func() {
		time.Sleep(time.Second)
		//c.AsyncWrite(data)
		clientArgs := &ClientArgs{}
		err := json.Unmarshal(data, clientArgs)
		if err != nil {
			c.AsyncWrite([]byte(fmt.Sprintf("json unmarshal error: %v", err)))
			return
		}
		handler := s.handlers[clientArgs.Handler]
		err = handler.Handle(c, clientArgs.Service, []byte(clientArgs.Model))
		if err != nil {
			c.AsyncWrite([]byte(fmt.Sprintf("[error] service: %s, err: %v", clientArgs.Service, err)))
		}
	}()
	return
}

type Server struct {
	opt       *Options
	tcpServer *TcpServer
}

type InitServicesMap func() map[string]Handler

func New(opt *Options, init InitServicesMap) (*Server, error) {
	tcpServer := &TcpServer{
		tick:      opt.BeatInterval * time.Second,
		port:      opt.Port,
		multicore: opt.MultiCore,
		handlers:  init(),
	}

	s := &Server{
		opt:       opt,
		tcpServer: tcpServer,
	}

	return s, nil
}

func (s *Server) Start() error {
	gameServer := s.tcpServer

	addr := fmt.Sprintf("tcp://:%d", gameServer.port)

	if err := gnet.Serve(gameServer, addr, gnet.WithMulticore(gameServer.multicore), gnet.WithTicker(gameServer.tick > 0)); err != nil {
		return errors.Wrap(err, "tcp server serve error")
	}
	return nil
}

var ProviderSet = wire.NewSet(New, NewOptions)
