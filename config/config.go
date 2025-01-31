package config

import "os"

type Config struct {
	CSVFilePath string
}

func NewConfig() *Config {
	csvFilePath := os.Getenv("CSV_FILE_PATH")
	if csvFilePath == "" {
		csvFilePath = "./data/users.csv"
	}

	return &Config{CSVFilePath: csvFilePath}

}
