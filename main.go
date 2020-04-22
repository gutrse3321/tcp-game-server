package main

import (
	"flag"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/20 11:08
 * @Title:
 * --- --- ---
 * @Desc:
 */
var configFile = flag.String("f", "config/config.yml", "set config file which viper will loading")

func main() {
	flag.Parse()

	server, err := CreateServer(*configFile)
	if err != nil {
		panic(err)
	}

	if err := server.Start(); err != nil {
		panic(err)
	}
}
