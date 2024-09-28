package main

import "github.com/rizky201008/okx-wallet/app"

func main() {
	viper := app.InitViper()
	app.InitDb(viper)
	app.InitFiber(viper)
}
