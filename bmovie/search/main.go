package main

import (
	"flag"
	"github.com/danClauz/bibit/bmovie/search/di"
	"github.com/danClauz/bibit/bmovie/search/infrastructure"
	"github.com/danClauz/bibit/bmovie/search/shared"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	flag.Parse()
	container := di.Container

	err := container.Invoke(func(sh shared.Holder, inh infrastructure.Holder) error {
		sh.Logger.Println("running a rest api application")

		sig := make(chan os.Signal)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

		go inh.ServeHttp()
		go inh.ServeGateway()
		go inh.ServeGrpc()

		sh.Logger.Println("receive signal", <-sig)
		sh.Close()

		return nil
	})

	if err != nil {
		panic(err)
	}
}
