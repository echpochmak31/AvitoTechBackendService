package configs

import (
	"errors"
	"os"
)

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DbName   string
}

func GetPostgresUrl() (string, error) {
	cfg, err := GetDatabaseConfig()
	if err != nil {
		return "", err
	}
	dbUrl := "postgres://" + cfg.Username + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port + "/" + cfg.DbName
	return dbUrl, nil
}

func GetDatabaseConfig() (DatabaseConfig, error) {
	cfg := DatabaseConfig{}
	username, err := getEnvVariable("POSTGRES_USER")
	if err != nil {
		return DatabaseConfig{}, err
	}
	password, err := getEnvVariable("POSTGRES_PASSWORD")
	if err != nil {
		return DatabaseConfig{}, err
	}
	host, err := getEnvVariable("DB_HOST")
	if err != nil {
		return DatabaseConfig{}, err
	}
	port, err := getEnvVariable("POSTGRES_PORT_DST")
	if err != nil {
		return DatabaseConfig{}, err
	}
	dbName, err := getEnvVariable("DB_NAME")
	if err != nil {
		return DatabaseConfig{}, err
	}
	cfg.Username = username
	cfg.Password = password
	cfg.Host = host
	cfg.Port = port
	cfg.DbName = dbName
	return cfg, nil
}

func getEnvVariable(varName string) (string, error) {
	if value, exists := os.LookupEnv(varName); exists {
		return value, nil
	}
	return "", errors.New("Failed to get env variable: " + varName)
}
