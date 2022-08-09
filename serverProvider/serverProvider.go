package serverProvider

import (
	"example.com/m/provider"
	kafkaprovider "example.com/m/provider/kafkaProvider"
	"example.com/m/provider/realtimesocketmanager"
	"net/http"
	"time"

	"context"

	"github.com/sirupsen/logrus"
)

type Server struct {
	//StorageProvider provider.StorageProvider
	httpServer  *http.Server
	RealtimeHub provider.WebSocketHubProvider
	Messenger   provider.KafkaProvider
}

func SrvInit() *Server {
	messenger := kafkaprovider.NewKafkaProvider()

	realtimeHub := realtimesocketmanager.NewRealtimeHub(messenger)

	//sp := provider.NewStorageProvider(os.Getenv("GCP_Key"))
	return &Server{
		//StorageProvider: sp,
		RealtimeHub: realtimeHub,
		Messenger:   messenger,
	}
}

func (srv *Server) Start() {
	//addr := ":" + os.Getenv("PORT")
	addr := ":4444"
	httpSrv := &http.Server{
		Addr:    addr,
		Handler: srv.SetupRoutes(),
	}

	srv.httpServer = httpSrv

	go srv.RealtimeHub.Run()

	logrus.Info("Server running at PORT ", addr)
	if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.Fatalf("Start %v", err)
		return
	}
}

func (srv *Server) Stop() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logrus.Info("closing server...")
	_ = srv.httpServer.Shutdown(ctx)
	logrus.Info("Done")
}
