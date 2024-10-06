package application

import (
	"os"
	"strconv"
)

type Config struct {
	RedisAddress string
	ServerPort   uint16
}

func LoadConfig() Config {
	cfg := Config{
		RedisAddress: "localhost:6379",
		ServerPort:   3000,
	}

	if reddisAddr, exists := os.LookupEnv("REDIS_ADDRESS"); exists {
		cfg.RedisAddress = reddisAddr
	}

	if serverPort, exists := os.LookupEnv("SERVER_PORT"); exists {
		if port, err := strconv.ParseUint((serverPort), 10, 16); err == nil {
			cfg.ServerPort = uint16(port)
		}
	}

	return cfg
}
