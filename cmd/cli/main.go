package main

import (
	"github.com/spf13/cobra"
	"github.com/wokacz/mamba/internal/env"
	"github.com/wokacz/mamba/internal/files"
)

type config struct {
	name string
}

type application struct {
	config *config
	cobra  *cobra.Command
}

var app = &application{}
var f = &files.Files{}

func init() {
	app.config = &config{}
	app.config.name = env.GetString("APP_NAME", "mamba")
	app.cobra = &cobra.Command{
		Use:   app.config.name,
		Short: "CLI to generate FastAPI projects",
	}
}

func main() {
	app.cobra.AddCommand(app.initCommand())
	app.cobra.AddCommand(app.doctorCommand())

	err := app.cobra.Execute()
	if err != nil {
		panic(err)
	}
}
