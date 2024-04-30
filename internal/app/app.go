package app

import (
	"carsRegistry/internal/config"
	"carsRegistry/internal/delivery"
	"carsRegistry/internal/integration"
	"carsRegistry/internal/repository"
	"carsRegistry/internal/service"
	l "carsRegistry/pkg/logg"
	"fmt"
	lr "github.com/sirupsen/logrus"
	"net/http"
)

func Run() {
	l.LogInfo("app:Run", "Starting the application", lr.Fields{})

	cfg, err := config.LoadConfig(".env")
	if err != nil {
		l.LogError("app:Run", "Failed to load config", err, lr.Fields{})
		return
	}
	l.LogInfo("app:Run", "Config loaded", lr.Fields{})

	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
	db, err := repository.NewPostgresDB(connString)
	if err != nil {
		l.LogError("app:Run", "Failed to connect to the database", err, lr.Fields{"connString": connString})
		return
	}
	l.LogInfo("app:Run", "Connected to the database", lr.Fields{})

	depIntegrations := integration.DepIntegrations{CarsInfoURL: cfg.CarsInfoURL}

	repos := repository.NewRepository(db)
	integrations := integration.NewIntegrations(depIntegrations)
	services := service.NewServices(repos, integrations)
	handler := delivery.NewHandler(services)

	var srv http.Server
	srv.Handler = handler.InitRoutes()
	srv.Addr = cfg.SRVHost + ":" + cfg.SRVPort
	l.LogInfo("app:Run", "Server started", lr.Fields{"host": cfg.SRVHost, "port": cfg.SRVPort})
	srv.ListenAndServe()
	l.LogInfo("app:Run", "Server stopped", lr.Fields{})
}
