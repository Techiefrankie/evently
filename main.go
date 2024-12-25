package main

import (
	"evently/config"
)

func main() {
	config.InitGorm()
	RegisterRoutes()
}

type App struct {
	Env map[string]string
}

func GetApp() *App {
	envs := make(map[string]string)

	envs["PORT"] = config.GetEnv(config.Port, "8080")
	envs["SECRET_KEY"] = config.GetEnv(config.SecretKey, "")

	return &App{
		Env: envs,
	}
}
