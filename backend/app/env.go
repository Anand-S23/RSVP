package app

import (
	"fmt"
	"log"
	"os"
)

type EnvVars struct {
    PRODUCTION           bool
    PORT                 string
    FE_URI               string
    DB_URI               string
}

func LoadEnv() (*EnvVars) {
    envMode       := MustGetEnv("MODE")
    port          := MustGetEnv("PORT")
    feURI         := MustGetEnv("FE_URI")
    dbURI         := MustGetEnv("DB_URI")

    return &EnvVars {
        PRODUCTION: (envMode == "production"),
        FE_URI: feURI,
        DB_URI: dbURI,
        PORT: port,
    }
}

func MustGetEnv(env string) string {
	variable := os.Getenv(env)
	if variable == "" {
        message := fmt.Sprintf("Must provide %s variable in .env file", env)
        log.Fatal(message)
	}

	return variable
} 

