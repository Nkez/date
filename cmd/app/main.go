package main

import (
	"github.com/Nkez/date/internal/handler"
	"github.com/Nkez/date/internal/interfaces/geolite"
	"github.com/Nkez/date/internal/interfaces/postgres"
	"github.com/Nkez/date/internal/repositories"
	"github.com/Nkez/date/internal/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
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
	if err := fasthttp.ListenAndServe(":8000", handlers.InitRouter().HandleRequest); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func ConfigInit() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

//func timeFunc(context *routing.Context) error {
//	//db, err := geoip2.Open("GeoLite2-Country.mmdb")
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	db2, err := geoip2.Open("GeoLite2-City.mmdb")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Fprintf(context, "Hi there! RequestURI is %q", context.RequestURI())
//	clientIP := realip.FromRequest(context.RequestCtx)
//	fmt.Print(clientIP)
//	ua := string(context.Request.Header.UserAgent())
//	parsed := user_agent.New(ua)
//	f, _ := parsed.Browser()
//	c := parsed.OS()
//	g := parsed.UA()
//	m := parsed.Mobile()
//	city, _ := db2.City(net.ParseIP("188.169.10.208"))
//	loc, _ := time.LoadLocation(fmt.Sprintf("Asia/Tbilisi"))
//	fmt.Println(city.Continent.Names["en"])
//	fmt.Println(city.City.Names["en"])
//	now := time.Now().In(loc)
//	//country, _ := db.Country(net.ParseIP("178.236.62.197"))
//	fmt.Fprintf(context, "Hi there! IP is %q, %q, %q, %q,%q,%q,%q", clientIP, f, c, g, m, city.Country.Names["en"], city.City.Names["en"])
//	fmt.Fprintf(context, "Hi there!%q", now)
//	//context.Redirect("https://api.ipbase.com/v2/info?ip="+clientIP+"&apikey=u5tRuvmJLL6x44gWVv7c5CaOgEoFUMNoBQ4z0kLr", http.StatusSeeOther)
//	context.SetContentType("text/plain; charset=utf8")
//
//	// Set arbitrary headers
//	context.Response.Header.Set("X-My-Header", "my-header-value")
//	return nil
//}
//
//func dateFunc(context *routing.Context) error {
//	fmt.Fprintf(context, "Hi there! RequestURI is %q", context.RequestURI())
//	return nil
//}

//
////func barHandlerFunc(ctx *fasthttp.RequestCtx) {
////	fmt.Fprintf(ctx, "Hi there! RequestURI is %q", ctx.RequestURI())
////}
////
////func fooHandlerFunc(ctx *fasthttp.RequestCtx) {
////	fmt.Fprintf(ctx, "Hi there! RequestURI is %q", ctx.RequestURI())
////}
