package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/panjf2000/gnet"
	"log"
	"sync"
	"time"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/20 11:08
 * @Title:
 * --- --- ---
 * @Desc:
 */
type echoServer struct {
	*gnet.EventServer
	tick             time.Duration
	connectedSockets sync.Map
}

func (es *echoServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	log.Printf("Echo server is listening on %s (multi-cores: %t, loops: %d)\n",
		srv.Addr.String(), srv.Multicore, srv.NumEventLoop)
	return
}

func (es *echoServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	log.Printf("Socket with addr: %s has been opened...\n", c.RemoteAddr().String())
	es.connectedSockets.Store(c.RemoteAddr().String(), c)
	return
}
func (es *echoServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	log.Printf("Socket with addr: %s is closing...\n", c.RemoteAddr().String())
	es.connectedSockets.Delete(c.RemoteAddr().String())
	return
}

func (es *echoServer) Tick() (delay time.Duration, action gnet.Action) {
	log.Println("heart beat ... ")
	es.connectedSockets.Range(func(key, value interface{}) bool {
		addr := key.(string)
		c := value.(gnet.Conn)
		c.AsyncWrite([]byte(fmt.Sprintf("heart beating to %s\n", addr)))
		return true
	})
	delay = es.tick
	return
}

type ClientArgs struct {
	Service string
	Model   string
}

type Model struct {
	NickName string
}

func (es *echoServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	// Echo synchronously.
	out = frame
	clientArgs := &ClientArgs{}
	err := json.Unmarshal(frame, clientArgs)
	if err != nil {
		log.Println("err:", err)
	}
	model := &Model{}
	json.Unmarshal([]byte(clientArgs.Model), model)
	log.Println("client:", model)
	//log.Println("client:", string(frame))
	return

	/*
		// Echo asynchronously.
		data := append([]byte{}, frame...)
		go func() {
			time.Sleep(time.Second)
			c.AsyncWrite(data)
		}()
		return
	*/
}

func main() {
	var port int
	var multicore bool
	var interval time.Duration
	var ticker bool

	// Example command: go run push.go --port 9000 --tick 1s --multicore=true
	flag.IntVar(&port, "port", 9000, "server port")
	flag.BoolVar(&multicore, "multicore", true, "multicore")
	flag.DurationVar(&interval, "tick", time.Second*10, "pushing tick")
	flag.Parse()
	if interval > 0 {
		ticker = true
	}
	push := &echoServer{tick: interval}
	log.Fatal(gnet.Serve(push, fmt.Sprintf("tcp://:%d", port), gnet.WithMulticore(multicore), gnet.WithTicker(ticker)))
}
