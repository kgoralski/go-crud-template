package servid

import (
	"flag"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/kgoralski/go-crud-template/internal/banks"
	"github.com/kgoralski/go-crud-template/internal/platform/db"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

const (
	defaultConfigFilePath  = "./configs"
	configFilePathUsage    = "config file directory. Config file must be named 'conf_{env}.yml'."
	configFilePathFlagName = "configFilePath"
	envUsage               = "environment for app, prod, dev, test"
	envDefault             = "prod"
	envFlagname            = "env"
)

var configFilePath string
var env string

func config() {
	logger()
	flag.StringVar(&configFilePath, configFilePathFlagName, defaultConfigFilePath, configFilePathUsage)
	flag.StringVar(&env, envFlagname, envDefault, envUsage)
	flag.Parse()
	configuration(configFilePath, env)
}

func logger() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
}

// App Instance which contains router and dao
type App struct {
	*http.Server
	r          *chi.Mux
	db         *sqlx.DB
	bankRouter *banks.Router
}

// NewApp creates new App with db connection pool
func NewApp() *App {
	config()
	router := chi.NewRouter()
	database := setupDB(viper.GetString("database.URL"))
	banksRouter := banks.NewRouter(router, database)
	server := &App{
		r:          router,
		db:         database,
		bankRouter: banksRouter,
	}
	server.routes()
	return server
}

// Start launching the server
func (a *App) Start() {
	log.Fatal(http.ListenAndServe(viper.GetString("server.port"), a.r))
}

func (a *App) routes() {
	a.bankRouter.Routes()
	showRoutes(a.r)
}

func showRoutes(r *chi.Mux) {
	log.Info("registered routes: ")
	walkFunc := func(method string, route string, handler http.Handler, m ...func(http.Handler) http.Handler) error {
		log.Infof("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(r, walkFunc); err != nil {
		log.Infof("Logging err: %s\n", err.Error())
	}
}

func configuration(path string, env string) {
	if flag.Lookup("test.v") != nil {
		env = "test"
		path = "./../../configs"
	}
	log.Println("Environment is: " + env + " configFilePath is: " + path)
	viper.SetConfigName("conf_" + env)
	viper.AddConfigPath(path) // working directory
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("fatal: %+v", err))
	}
}

func setupDB(dbURL string) *sqlx.DB {
	mysql, err := db.New(dbURL)
	if err != nil {
		log.Fatal(fmt.Errorf("fatal: %+v", err))
	}
	return mysql
}
