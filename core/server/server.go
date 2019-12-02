package coreserver

import (
	"net"
)

var serverConfig *ServerConfig

type ServerConfig struct {
	IpAddress string
	// UrlPath   string
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

func GetServerConfig() *ServerConfig {
	return serverConfig
}
