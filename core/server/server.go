package coreserver

import (
	"net"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

var serverConfig *ServerConfig

var appConfig *AppConfig

type ServerConfig struct {
	IpAddress string
}

type AppConfig struct {
	AppLogPath string
	RedisHost  string
	RedisPort  string
	RedisDB    int
}

func init() {
	godotenv.Load()
}

func SetServerGonfig() {

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				serverConfig = &ServerConfig{
					IpAddress: ipnet.IP.String(),
				}
			}
		}
	}
}

func SetAppConfig() {
	redisDb, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	appConfig = &AppConfig{
		AppLogPath: os.Getenv("APP_LOG_PATH"),
		RedisHost:  os.Getenv("REDIS_HOST"),
		RedisPort:  os.Getenv("REDIS_PORT"),
		RedisDB:    redisDb,
	}
}

func GetServerConfig() *ServerConfig {
	return serverConfig
}

func GetAppConfig() *AppConfig {
	return appConfig
}
