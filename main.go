package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"example.com/m/serverProvider"

	"github.com/sirupsen/logrus"
)

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("running1")
	srv := serverProvider.SrvInit()
	fmt.Println("running2")
	go srv.Start()
	fmt.Println("running3")
	go srv.Pub()
	fmt.Println("running4")
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	//})
	fmt.Println("running")
	<-done
	logrus.Info("Graceful shutdown")
	srv.Stop()

}
