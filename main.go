package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"

	"github.com/RangelReale/osin"
	log "github.com/Sirupsen/logrus"
	"github.com/gocraft/web"
	_ "github.com/lib/pq"
	"github.com/mikerjacobi/oauth/controllers"
	"github.com/mikerjacobi/oauth/models"
	"github.com/spf13/viper"
)

func main() {
	var (
		config     string
		configpath string
	)

	flag.StringVar(&config, "config", "config", "The base-name of your config file, extension may be .toml, .json, or .yml")
	flag.StringVar(&configpath, "configpath", ".", "the location of your config file")
	flag.Parse()

	viper.AddConfigPath(configpath)
	viper.AddConfigPath("/etc/app")
	viper.SetConfigName(config)
	viper.ReadInConfig()

	postgresHostVar := viper.GetStringMapString("postgres")["host_env_var"]
	viper.BindEnv(postgresHostVar)
	host := viper.Get(postgresHostVar)
	user := viper.GetStringMapString("postgres")["user"]
	log.Infof("dbhost: %s", host)
	conn := fmt.Sprintf("postgres://%s@%s/accounts?sslmode=disable", user, host)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Errorf("failed to connect")
		return
	}

	s := models.NewStorage(db)
	osinConfig := osin.NewServerConfig()
	osinConfig.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.CODE, osin.TOKEN}
	server := osin.NewServer(osinConfig, s)

	context := controllers.AuthContext{*server, db}
	router := web.New(context).
		Get("/authorize", context.AuthorizeHandler).
		Get("/token", context.TokenHandler).
		Get("/healthcheck", context.HealthCheckHandler)

	port := viper.GetString("server_port")
	log.Infof("Listening on port: %s", port)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), router)
}
