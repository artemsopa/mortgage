package main

import "github.com/artomsopun/mortgage/mortgage-api/internal/app"

const configsDir = "configs"

func main() {
	app.Run(configsDir)
}
