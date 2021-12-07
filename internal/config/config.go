package config

type Config struct {
	RedisDB    string `env:"REDIS_URL"`
	ServerPort int    `env:"SERVER_PORT" envDefault:"3000"`
}
