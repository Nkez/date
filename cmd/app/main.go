package main

import (
	"errors"
	date_protobuf "github.com/Nkez/date-protobuf"
	"github.com/Nkez/date/internal/handler"
	"github.com/Nkez/date/internal/interfaces/geolite"
	grpc2 "github.com/Nkez/date/internal/interfaces/grpc"
	"github.com/Nkez/date/internal/interfaces/postgres"
	"github.com/Nkez/date/internal/model"
	"github.com/Nkez/date/internal/repositories"
	"github.com/Nkez/date/internal/service"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
	"github.com/valyala/fasthttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

var configPath = ""

func main() {
	app := &cli.App{
		Name:  "date",
		Usage: "service",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "Load configuration from `FILE`",
				EnvVars:     []string{"DATE_CONFIG_PATH"},
				TakesFile:   true,
				Destination: &configPath,
				HasBeenSet:  false,
			},
		},
		Commands: []*cli.Command{
			{
				Name:      "migrate",
				Usage:     "Run migrations",
				Action:    runMigrate,
				ArgsUsage: "",
			},
		},
		Action: startApp,
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

func ConfigInit() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func startApp(context *cli.Context) error {
	if err := ConfigInit(); err != nil {
		logrus.Fatalf("error instaling configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	db, err := postgres.NewPostgresDB(&model.Config{
		DBUrl: viper.GetString("db"),
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
		err = serv.Serve(lis)
		if err != nil {
			return
		}
	}()
	if err := fasthttp.ListenAndServe(viper.GetString("port"), handlers.InitRouter().HandleRequest); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
	return nil
}

func runMigrate(cliContext *cli.Context) error {
	err := ConfigInit()
	if err != nil {
		return err
	}

	args := cliContext.Args()

	cfg := &model.Config{
		DBUrl: viper.GetString("db"),
	}
	if args.First() == "up" {
		if err := postgres.Up(cfg); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				return nil
			}
			return err
		}
	}
	if args.First() == "down" {
		if err := postgres.Down(cfg); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				return nil
			}
			return err
		}
	}

	return nil
}
