package config

//EmailConfig - email connection info
type EmailConfig struct {
	Email       string
	Password    string
	SMTPAddress string
	SMTPPort    int
}

//Config - runtime config
type Config struct {
	Email    EmailConfig
	EmailKey string
	Port     string
	Host     string
}
