package main

import (
	"permission-service/internal/app"
	"permission-service/internal/config/env"
)

func main() {
	if err := env.ValidateEnv(); err != nil {
		panic("error validating env: " + err.Error())
	}
	port := env.PORT
	app.New(port)
}
