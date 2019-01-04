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
)

const (
	defaultConfigFilePath  = "./configs"
	configFilePathUsage    = "config file directory. Config file must be named 'conf_{env}.yml'."
	configFilePathFlagName = "configFilePath"
	envUsage               = "environment for app, prod, dev, test"
	envDefault             = "dev"
	envFlagname            = "env"
)

var configFilePath string
var env string

func config() {
	flag.StringVar(&configFilePath, configFilePathFlagName, defaultConfigFilePath, configFilePathUsage)
	flag.StringVar(&env, envFlagname, envDefault, envUsage)
	flag.Parse()
	configuration(configFilePath, env)
}

// ServerInstance Instance which contains router and dao
type ServerInstance struct {
	*http.Server
	r          *chi.Mux
	db         *sqlx.DB
	bankRouter *banks.Router
}

// NewServer creates new ServerInstance with db connection pool
func NewServer() *ServerInstance {
	config()
	router := chi.NewRouter()
	database := setupDB(viper.GetString("database.URL"))
	banksRouter := banks.NewRouter(router, database)
	server := &ServerInstance{
		r:          router,
		db:         database,
		bankRouter: banksRouter,
	}
	server.routes()
	return server
}

// Start launching the server
func (s *ServerInstance) Start() {
	log.Fatal(http.ListenAndServe(viper.GetString("server.port"), s.r))
}

func (s *ServerInstance) routes() {
	s.bankRouter.Routes()
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
