package app

import (
	"DnDApi/internal/config"
	"DnDApi/internal/handler"
	"DnDApi/internal/repository"
	"DnDApi/internal/repository/postgres"
	"DnDApi/internal/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func Run() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	cfg, err := config.Init()
	if err != nil {
		logrus.Error(err)

		return
	}

	//зависимости
	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		Username: cfg.Postgres.Username,
		DBName:   cfg.Postgres.DBName,
		Password: cfg.Postgres.Password,
		SSLMode:  cfg.Postgres.SSLMode,
	})
	if err != nil {
		logrus.Fatalf("error start DB : %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(Server)
	if err = server.Run(cfg.Server.Port, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error start server : %s", err.Error())
	}

}
