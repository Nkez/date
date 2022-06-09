package main

import (
	date_protobuf "github.com/Nkez/date-protobuf"
	"github.com/Nkez/date/internal/handler"
	"github.com/Nkez/date/internal/interfaces/geolite"
	grpc2 "github.com/Nkez/date/internal/interfaces/grpc"
	"github.com/Nkez/date/internal/interfaces/postgres"
	"github.com/Nkez/date/internal/repositories"
	"github.com/Nkez/date/internal/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func main() {
	if err := ConfigInit(); err != nil {
		logrus.Fatalf("error instaling configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db : %s", err.Error())
	}
	geo := geolite.NewGeoDB(geolite.Config{
		DB: viper.GetString("geo.path"),
	})
	repos := repositories.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services, geo)
	g := grpc2.NewEventApiStruct(repos)
	go func() {
		lis, err := net.Listen("tcp", viper.GetString("grpc"))
		if err != nil {
			logrus.Fatal("err grpc server")
		}
		serv := grpc.NewServer()
		date_protobuf.RegisterEventServiceServer(serv, g)
		reflection.Register(serv)
		serv.Serve(lis)
	}()
	if err := fasthttp.ListenAndServe(viper.GetString("port"), handlers.InitRouter().HandleRequest); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func ConfigInit() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
