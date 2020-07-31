package router

import (
	"context"
	"github.com/zhaojiasanxing/go_gateway/conf"
	"log"
	"net/http"
	"time"
)

var (
	HttpSrvHandler *http.Server
)

func HttpServerRun()  {
	r := InitRouter()
	HttpSrvHandler = &http.Server{
		Addr:             conf.GlobalConfig.Server.Port,
		Handler:r,
	}
	go func() {
		log.Printf(" [INFO] HttpServerRun:%s\n",conf.GlobalConfig.Server.Port)
		if err := HttpSrvHandler.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", conf.GlobalConfig.Server.Port, err)
		}
	}()
}

func HttpServerStop()  {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	log.Printf(" [INFO] HttpServerStop stopped\n")
}


