package main

import (
	"net"
	"net/http"
	"strconv"

	"github.com/clzhan/SimpleHlsServer/httpserver"
	"github.com/clzhan/SimpleHlsServer/log"
	"github.com/clzhan/SimpleHlsServer/utils"
	"github.com/clzhan/SimpleHlsServer/conf"
)

//远程获取pprof数据
func InitPprof() {
	//获取本机ip
	//rtmpAddr := fmt.Sprintf("%s:%d", network.GetLocalIpAddress(),6399)
	//
	//str ,_ := network.IntranetIP()
	//log.Info("local ip: ",str)

	go func() {
		//http://10.10.6.162:6399/debug/pprof
		pprofAddress := util.GetLocalIp()
		pprofAddress += ":"
		pprofAddress += strconv.Itoa(6399)

		log.Info(http.ListenAndServe(pprofAddress, nil))
	}()

}

func startHttpServer() error {
	var httpServerListen net.Listener
	var err error

	HttpFlsAddress := util.GetLocalIp()
	HttpFlsAddress += ":"
	HttpFlsAddress += conf.AppConf.WebPort

	httpServerListen, err = net.Listen("tcp", HttpFlsAddress)

	if err != nil {
		log.Error(err)
		return err
	}

	httpServer := httpserver.NewHttpServer()
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Error("HTTP server panic: ", r)
			}
		}()
		log.Info("HttpServer listen On", HttpFlsAddress)
		httpServer.Serve(httpServerListen)
	}()
	return err
}


func main() {

	conf.Init()
	log.Init()

	InitPprof()

	err := startHttpServer()
	if err != nil{
		log.Info("ListenAndServerHttpServer error :", err)
	}

	//rtmp.ConnectPull("rtmp://10.10.6.39:1935/live/movie")
	//rtmp.ConnectPush("rtmp://10.10.6.39:1935/live/movie")
	log.Info("Start http Server.....")
	// do event loop
	select {}

}
