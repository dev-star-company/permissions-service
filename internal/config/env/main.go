package env

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var (
	PG_HOST     string
	PG_USER     string
	PG_PASSWORD string
	PG_DBNAME   string
	PG_SSLMODE  string
	PG_PORT     int
	PORT        int
	DEV_MODE    bool
)

var envFile map[string]string

func init() {
	var err error

	if strings.ToLower(os.Getenv("LAUNCH_MODE")) == "debug" {
		absPath, pathErr := filepath.Abs(os.Args[0] + "../../.env")
		if pathErr != nil {
			panic("error resolving absolute path: " + pathErr.Error())
		}

		envFile, err = godotenv.Read(absPath)
	} else {
		envFile, err = godotenv.Read(".env")
	}
	if err != nil {
		panic("error reading .env file: " + err.Error())
	}

	PG_HOST = EnvToString("PG_HOST")
	PG_USER = EnvToString("PG_USER")
	PG_PASSWORD = EnvToString("PG_PASSWORD")
	PG_DBNAME = EnvToString("PG_DBNAME")
	PG_SSLMODE = EnvToString("PG_SSLMODE")
	PG_PORT = EnvToInt("PG_PORT")
	PORT = EnvToInt("PORT")
	DEV_MODE = EnvToBool("DEV_MODE")
}

func init() {
	var err error
	fmt.Println("Execution path:", os.Args[0])
	if strings.ToLower(os.Getenv("LAUNCH_MODE")) == "debug" {
		envFile, err = godotenv.Read(".env")
	} else {
		envFile, err = godotenv.Read(".env")
	}
	if err != nil {
		panic("error reading .env file: " + err.Error())
	}
}

func ValidateEnv() error {
	requiredVars := []string{"PG_HOST", "PG_USER", "PG_PASSWORD", "PG_DBNAME", "PG_PORT", "PG_SSLMODE", "PORT", "DEV_MODE"}
	for _, key := range requiredVars {
		if envFile[key] == "" {
			return errors.New(key + " is not set in .env file")
		}
	}
	return nil
}

func EnvToString(v string) string {
	return envFile[v]
}

func EnvToBool(v string) bool {
	return strings.ToLower(envFile[v]) == "true"
}

func EnvToInt(v string) int {
	vInt, err := strconv.Atoi(envFile[v])
	if err != nil {
		panic("error converting " + v + " to int: " + err.Error())
	}
	return vInt
}
