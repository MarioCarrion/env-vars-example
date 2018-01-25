// Simple example for providing default environment variables via an env file
// and wrapping those into a struct with the possibility to override them
// if needed
package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type DatabaseSourceName struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var dsn DatabaseSourceName
	err = envconfig.Process("dsn", &dsn)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = fmt.Printf("%s:%s@%s:%d/%s\n", dsn.Username, dsn.Password, dsn.Host, dsn.Port, dsn.Database)
	if err != nil {
		log.Fatal(err.Error())
	}
}
