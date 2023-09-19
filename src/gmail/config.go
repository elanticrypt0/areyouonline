package gmail

type Config struct {
	Email  string
	Passwd string
}

func GetConfig() *Config {
	return &Config{
		Email:  "golaschgo@gmail.com",
		Passwd: "Remorse-Quintet-Status-Squad-Hunk2-Manned",
	}
}
