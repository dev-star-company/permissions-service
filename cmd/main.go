package main

import (
	"permissions-service/internal/app"
	_ "permissions-service/internal/app/ent/runtime"
	"permissions-service/internal/config/env"
)

func main() {
	if err := env.ValidateEnv(); err != nil {
		panic("error validating env: " + err.Error())
	}
	port := env.PORT
	app.New(port)
}
