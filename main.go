package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/dailoi280702/vrs-general-service/client/mysql"
	_ "github.com/dailoi280702/vrs-general-service/client/redis"
	"github.com/dailoi280702/vrs-general-service/config"
	"github.com/dailoi280702/vrs-general-service/handler/grpc"
	"github.com/dailoi280702/vrs-general-service/log"
	"github.com/dailoi280702/vrs-general-service/migration"
	"github.com/dailoi280702/vrs-general-service/proto"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.GetConfig()

	var (
		logger = log.Logger()
		errs   = make(chan error)
		g      = grpc.NewGRPCServer()
	)

	if err := migration.Migrate(mysql.GetClient()); err != nil {
		logger.Error("Migration failed", "error", err)
	}

	defer g.GracefulStop()

	go func() {
		l, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", cfg.Port))
		if err != nil {
			logger.Error("Failed to listen", "port", cfg.Port, "error", err)
		}

		if cfg.IsDebug {
			reflection.Register(g)
		}

		proto.RegisterServiceServer(g, &grpc.Service{})
		errs <- g.Serve(l)
	}()

	go func() {
		for err := range errs {
			logger.Error("Server error", "error", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-c

	logger.Info("The server is stopping ...")

	close(c)
	close(errs)
}
