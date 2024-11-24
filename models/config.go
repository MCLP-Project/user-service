package models

type Config struct {
	ServerConfig   ServerConfig   `json:"serverConfig"`
	DatabaseConfig DatabaseConfig `json:"databaseConfig"`
}
type ServerConfig struct {
	Port string `json:"port"`
}
type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
