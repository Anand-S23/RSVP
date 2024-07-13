package main

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
    JWT_SECRET           []byte
    COOKIE_HASH_KEY      []byte
    COOKIE_BLOCK_KEY     []byte
    ADMIN_EMAIL          string
    ADMIN_PASSWORD       string
}

func LoadEnv() (*EnvVars) {
    envMode       := MustGetEnv("MODE")
    port          := MustGetEnv("PORT")
    feURI         := MustGetEnv("FE_URI")
    dbURI         := MustGetEnv("DB_URI")
    secret        := MustGetEnv("JWT_SECRET")
    hashKey       := MustGetEnv("COOKIE_HASH_KEY")
    blockKey      := MustGetEnv("COOKIE_BLOCK_KEY")
    adminEmail    := MustGetEnv("ADMIN_EMAIL")
    adminPassword := MustGetEnv("ADMIN_PASSWORD")

    return &EnvVars {
        PRODUCTION: (envMode == "production"),
        FE_URI: feURI,
        DB_URI: dbURI,
        JWT_SECRET: []byte(secret),
        COOKIE_HASH_KEY: []byte(hashKey),
        COOKIE_BLOCK_KEY: []byte(blockKey),
        PORT: port,
        ADMIN_EMAIL: adminEmail,
        ADMIN_PASSWORD: adminPassword,
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

