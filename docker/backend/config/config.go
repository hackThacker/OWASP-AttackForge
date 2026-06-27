package config

import "os"

type Config struct {
	Port     string
	GRPCPort string
	Domain   string
}

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8081"
	} else if port[0] != ':' {
		port = ":" + port
	}

	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = ":50051"
	} else if grpcPort[0] != ':' {
		grpcPort = ":" + grpcPort
	}

	domain := os.Getenv("DOMAIN")
	if domain == "" {
		domain = "hackthacker.lab"
	}

	return &Config{
		Port:     port,
		GRPCPort: grpcPort,
		Domain:   domain,
	}
}
