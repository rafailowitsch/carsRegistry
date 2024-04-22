package app

import (
	"carRegistry/internal/config"
	"carRegistry/internal/delivery"
	"carRegistry/internal/repository"
	"carRegistry/internal/service"
	"fmt"
	"log"
	"net/http"
)

func Run() {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("Unable to load config: %v", err)
	}

	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := repository.NewPostgresDB(connString)
	if err != nil {
		log.Fatalf("Unable to connect to DB: %v", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewServices(repos)
	handler := delivery.NewHandler(services)

	var srv http.Server
	srv.Handler = handler.InitRoutes()
	srv.Addr = cfg.SRVHost + ":" + cfg.SRVPort
	log.Printf("Server started on %s", srv.Addr)
	srv.ListenAndServe()
}
