package config

import "os"

type Environment struct {
	ClickhouseHost     string
	ClickhouseDB       string
	ClickhouseUsername string
	ClickhousePass     string
	ServerHost         string
}

func New() Environment {
	return Environment{
		ClickhouseHost:     os.Getenv("CLICKHOUSEHOST"),
		ClickhouseDB:       os.Getenv("CLICKHOUSEDB"),
		ClickhouseUsername: os.Getenv("CLICKHOUSEUSERNAME"),
		ClickhousePass:     os.Getenv("CLICKHOUSEPASS"),
		ServerHost:         ":3000",
	}
}
