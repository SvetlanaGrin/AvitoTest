package main

import (
	"context"
	"github.com/SvetlanaGrin/AvitoTest"
	"github.com/SvetlanaGrin/AvitoTest/pkg/controller"
	"github.com/SvetlanaGrin/AvitoTest/pkg/repository"
	"github.com/SvetlanaGrin/AvitoTest/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(filepath.Join("C://Users/user/GolandProjects/AvitoTest/", ".env")); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Post:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	controllers := controller.NewControllers(services)

	srv := new(AvitoTest.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), controllers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}

}

func InitConfig() error {
	viper.AddConfigPath("C://Users/user/GolandProjects/AvitoTest/configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	return viper.ReadInConfig()
}
