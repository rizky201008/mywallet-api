package main

import "github.com/rizky201008/mywallet-backend/app"

func main() {
	viper := app.InitViper()
	app.InitDb(viper)
	app.InitFiber(viper)
}
