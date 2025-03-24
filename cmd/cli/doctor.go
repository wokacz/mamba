package main

import "github.com/spf13/cobra"

func (app *application) doctorCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "doctor",
		Short: "Check if all dependencies are installed",
		Run: func(cmd *cobra.Command, args []string) {
			app.checkPoetry()
		},
	}

	return cmd
}
