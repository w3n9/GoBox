package main

import (
	"GoBox/config"
	"GoBox/router"
)

func main(){
	webConf:=config.Get().Web
	addr:=webConf.Host+":"+webConf.Port
	err := router.Get().Run(addr)
	if err != nil {
		panic(err)
	}
}