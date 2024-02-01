package main

import (
	"chess/internal/app"
	"flag"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "config path")
	flag.Parse()

	if err := app.Run(configPath); err != nil {
		panic(err)
	}
}
