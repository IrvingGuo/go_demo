package server

import (
	"context"
	"go.uber.org/zap"
	"net/http"
	"resource-plan-improvement/config"
)

var log = config.Logger
var c = config.Conf

func NewServer(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("Panicking...", zap.Any("err", err))
		}
	}()

	server := &http.Server{
		Addr:    c.System.Addr,
		Handler: newRouter(),
	}

	go func() {
		log.Infof("Launching web server at %s", server.Addr)

		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				log.Info("Web server shutdown completely")
			} else {
				log.Error("Web server closed with exceptions", zap.Error(err))
			}
		}
	}()

	<-ctx.Done()
	log.Info("http: shutting down web server")
	err := server.Close()
	if err != nil {
		log.Error("Fail to shutdown the server", zap.Error(err))
	}
}
