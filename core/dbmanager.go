package core

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseConfig struct {
	ConnectionName string
	Engine         string
	Host           string
	Port           string
	User           string
	Password       string
	Dbname         string
}

func loadDBConfig() DatabaseConfig {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return DatabaseConfig{
		ConnectionName: os.Getenv("DB_CONNECTION_NAME"),
		Engine:         os.Getenv("DB_CONNECTION_TYPE"),
		Host:           os.Getenv("DB_HOST"),
		Port:           os.Getenv("DB_PORT"),
		Dbname:         os.Getenv("DB_NAME"),
		User:           os.Getenv("DB_USER"),
		Password:       os.Getenv("DB_PASSWORD"),
	}
}

func Connect2EngineDB(app_config *AppConfig) *gorm.DB {

	var db_connection *gorm.DB

	switch app_config.Db_config.Engine {
	case "sqlite":
		db_connection = DbConnectSqlite(app_config.Db_config.Dbname)
	case "postgres":
		db_connection = DbConnectPostgres(app_config.Db_config.Host, app_config.Db_config.User, app_config.Db_config.Password, app_config.Db_config.Dbname, app_config.Db_config.Port)
	case "mysql":
		db_connection = DbConnectMySql(app_config.Db_config.Host, app_config.Db_config.User, app_config.Db_config.Password, app_config.Db_config.Dbname, app_config.Db_config.Port)
	default:
		db_connection = DbConnectSqlite(app_config.Db_config.Dbname)
	}

	return db_connection
}

// Engines:

func DbConnectPostgres(host, user, password, dbname, port string) *gorm.DB {

	const dns = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable "
	// connect to gorn
	conn, err := gorm.Open(postgres.Open(fmt.Sprintf(dns, host, user, password, dbname, port)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect database Postgres")
	}
	return conn
}

func DbConnectMySql(host, user, password, dbname, port string) *gorm.DB {

	const dns = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dns_config := fmt.Sprintf(dns, user, password, host, port, dbname)
	// connect to gorn
	conn, err := gorm.Open(mysql.Open(dns_config), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect database MySQL")
	}
	return conn
}

// obtiene una conexión hacia la base de datos sqlite que se encuentra en la carpeta "libs" con el nombre que se le pase como parámetro
func DbConnectSqlite(dbname string) *gorm.DB {
	dbname = strings.ToLower(dbname)
	// gorm create sqlite db
	conn, err := gorm.Open(sqlite.Open("_db/"+dbname+".db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database SQLITE")
	}
	return conn
}
