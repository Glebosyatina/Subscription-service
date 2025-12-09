package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"glebosyatina/test_project/internal/handlers"
	"glebosyatina/test_project/internal/repository"
	"glebosyatina/test_project/internal/service"
	"glebosyatina/test_project/internal/service/user"
	"glebosyatina/test_project/pkg/database"
	"glebosyatina/test_project/server"
)

func main() {

	//читаем переменные из .env файла
	if err := godotenv.Load(); err != nil {
		log.Println("Error while reading variables from .env file")
	}

	//инициализация логгера
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	//инициализация репозиториев (уровень бд)
	db, err := database.NewDB(database.Config{
		Host:    os.Getenv("POSTGRES_HOST"),
		Port:    os.Getenv("POSTGRES_PORT"),
		User:    os.Getenv("POSTGRES_USER"),
		Passwd:  os.Getenv("POSTGRES_PASSWORD"),
		DBName:  os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("DB_SSL"),
	}, logger)
	if err != nil {
		logger.Error("Не удалось создать конфиг бд", err.Error())
		os.Exit(1)
	}
	//миграции
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.Error("Ошибка при создании драйвера для миграций", err.Error())
		os.Exit(1)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		logger.Error("Ошибка при создании инстанса бд для миграций")
		os.Exit(1)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logger.Error("Ошибка при миграции", slog.Any("error", err))
		os.Exit(1)
	}
	defer m.Close()

	defer func() {
		if err := db.Close(); err != nil {
			logger.Error("Ошибка при закрытии соединения с бд", err)
		} else {
			logger.Info("Соединение с бд успешно закрыто")
		}
	}()

	userRepo := repository.NewUserRepo(db)

	//инициализация сервисов (уровень бизнесс логики)

	services := &service.Services{
		UserService: user.NewUserService(userRepo, logger),
	}

	//инициализация мультиплексора(в аргументах передаем сервисы и логгер)
	mux := handlers.NewHandler(services, logger)

	//для gracefull shutdown отловим sigint sigterm
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	//создание и запуск сервера
	srv := new(server.Server)
	errorServer := make(chan error, 1)
	go func() {
		if err := srv.Run(os.Getenv("SERVER_PORT"), mux.InitRoutes()); err != nil {
			errorServer <- err
		}
	}()

	select {
	case <-c:
		logger.Info("Gracefull shutdown")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			logger.Error("Some error on server shutting down")
		}
		//block while context is done
		<-ctx.Done()
		logger.Info("Timeout is done")

		logger.Info("Server stopeed gracefully")

	case err := <-errorServer:
		logger.Error("Some error while starting server...", slog.Any("error", err))
		os.Exit(1)
	}

}
