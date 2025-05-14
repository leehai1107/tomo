package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

var (
	server   ServerCfg
	dbCfg    DBCfg
	services ServicesCfg
	cors     CorsCfg
)

type DBCfg struct {
	PgHost            string `envconfig:"PG_HOST" default:"db"`
	PgPort            string `envconfig:"PG_PORT" default:"5432"`
	PgUser            string `envconfig:"PG_USER" default:"postgres"`
	PgPassword        string `envconfig:"PG_PASSWORD" default:"mysecretpassword"`
	PgDatabase        string `envconfig:"PG_DATABASE" default:"postgres"`
	PgPoolSize        int    `envconfig:"PG_POOL_SIZE" default:"0"`
	PgIdleConnTimeout int    `envconfig:"PG_IDLE_CONNECTION_TIMEOUT" default:"30"`
	PgMaxConnAge      int    `envconfig:"PG_MAX_CONNECTION_AGE" default:"3000"`
	MongoURI          string `envconfig:"MONGO_URI" default:"0.0.0.0"`
}

type ServerCfg struct {
	ENV            string `envconfig:"ENVIRONMENT" default:"development"`
	SERVERUrl      string `envconfig:"SERVER_URL" default:"0.0.0.0"`
	GRPCPort       int    `envconfig:"USER_GRPC_PORT" default:"10000"`
	HTTPPort       int    `envconfig:"USER_HTTP_PORT" default:"8080"`
	LogLevel       string `envconfig:"LOG_LEVEL" default:"debug"`
	Production     bool   `envconfig:"PRODUCTION" default:"false"`
	GinMode        string `envconfig:"GIN_MODE" default:"debug"`
	Logger         bool   `envconfig:"LOGGER" default:"false"`
	CorsProduction bool   `envconfig:"CORS_PRODUCTION" default:"false"`
}

type ServicesCfg struct{}

type CorsCfg struct {
	Google   string `envconfig:"GOOGLE" default:"https://www.google.com/"`
	Facebook string `envconfig:"FACEBOOK" default:"https://www.facebook.com/"`
}

func InitConfig() {
	configs := []interface{}{
		&server,
		&services,
		&dbCfg,
		&cors,
	}
	for _, instance := range configs {
		err := envconfig.Process("", instance)
		if err != nil {
			log.Fatalf("unable to init config: %v, err: %v", instance, err)
		}
	}
}

func ServerConfig() ServerCfg {
	return server
}

func ServiceConfig() ServicesCfg {
	return services
}

func DBConfig() DBCfg {
	return dbCfg
}

func CorsConfig() CorsCfg {
	return cors
}
