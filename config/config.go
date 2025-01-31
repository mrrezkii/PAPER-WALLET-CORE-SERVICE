package config

import "os"

type Config struct {
	ServerPort  string
	CSVFilePath string
}

func NewConfig() *Config {
	csvFilePath := os.Getenv("CSV_FILE_PATH")
	if csvFilePath == "" {
		csvFilePath = "./data/users.csv"
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	return &Config{
		ServerPort:  serverPort,
		CSVFilePath: csvFilePath,
	}

}
