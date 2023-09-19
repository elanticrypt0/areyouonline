package core

type AppConfig struct {
	App_server_host   string
	App_server_port   string
	App_url           string
	App_setup_enabled bool
	App_debug_mode    bool
	App_CORS_Origins  string
	APP_CORS_Headers  string
	Db_config         DatabaseConfig
}
