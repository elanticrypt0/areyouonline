package core

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// Load the .env vars and return a struct
func LoadConfig(with_serverals_db bool) AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app_config := AppConfig{
		App_server_host:   os.Getenv("APP_SERVER_HOST"),
		App_server_port:   os.Getenv("APP_SERVER_PORT"),
		App_url:           os.Getenv("APP_SERVER_HOST") + ":" + os.Getenv("APP_SERVER_PORT"),
		App_setup_enabled: os.Getenv("APP_SETUP_ENABLED") == "true",
		App_debug_mode:    os.Getenv("APP_DEBUG_MODE") == "true",
		App_CORS_Origins:  os.Getenv("APP_CORS_ORIGINS"),
		APP_CORS_Headers:  os.Getenv("APP_CORS_HEADERS"),
	}

	if with_serverals_db == true {
		// TODO: carga la configuración de diferentes bases de datos.
	} else {
		app_config.Db_config = loadDBConfig()
	}

	return app_config

}

// TODO: Refactorizar para poder manegar varias conexiones a diferentes bases de datos dentro de una misma APP.
func Connect2DB(app_config *AppConfig) *gorm.DB {
	db_connection := DbConnectPostgres(app_config.Db_config.Host, app_config.Db_config.User, app_config.Db_config.Password, app_config.Db_config.Dbname, app_config.Db_config.Port)

	return db_connection
}

func Connect2OneDB(connection_name string, app_config *AppConfig) *gorm.DB {
	return Connect2EngineDB(app_config)
}
