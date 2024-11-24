package models

type Config struct {
	ServerConfig ServerConfig `json:"serverConfig"`
}
type ServerConfig struct {
	Port string `json:"port"`
}
