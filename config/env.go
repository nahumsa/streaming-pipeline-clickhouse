package config

import "os"

type Environment struct {
	ClickhouseHost     string
	ClickhouseDB       string
	ClickhouseUsername string
	ClickhousePass     string
	ServerHost         string
}

func getEnvVar(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic("clickhouseHost not in environment variables")
	}

	return v
}

func New() Environment {
	return Environment{
		ClickhouseHost:     getEnvVar("CLICKHOUSEHOST"),
		ClickhouseDB:       getEnvVar("CLICKHOUSEDB"),
		ClickhouseUsername: getEnvVar("CLICKHOUSEUSERNAME"),
		ClickhousePass:     getEnvVar("CLICKHOUSEPASS"),
		ServerHost:         ":3000",
	}
}
