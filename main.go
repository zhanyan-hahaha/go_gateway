package  main

import (
	"github.com/zhaojiasanxing/go_gateway/conf"
	"github.com/zhaojiasanxing/go_gateway/router"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main()  {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	runtime.GOMAXPROCS(runtime.NumCPU())

	router.HttpServerRun()
	defer conf.DB.Close()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}