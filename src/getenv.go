package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func getEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("env file missing")
	}

	return os.Getenv(key)
}
