package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wokacz/mamba/internal/resource"
	"github.com/wokacz/mamba/internal/words"
	"os"
	"time"
)

// doctorCommand checks if all dependencies are installed
func (a *application) doctorCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "doctor",
		Short: "Check if all dependencies are installed",
		Run: func(cmd *cobra.Command, args []string) {
			a.checkPython()
			a.checkPoetry()
		},
	}

	return cmd
}

// newCommand initializes a new FastAPI project
func (a *application) newCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new",
		Short: "Initialize a new FastAPI project",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var err error

			var projectName string
			if len(args) > 0 {
				projectName = args[0]
			} else {
				fmt.Print("Enter project name: ")
				reader := bufio.NewReader(os.Stdin)
				projectName, _ = reader.ReadString('\n')
				projectName = projectName[:len(projectName)-1] // Remove newline character

				// Start the loader in a separate goroutine
				done := make(chan bool)
				go loader(done)

				// Stop the loader
				defer func() { done <- true }()
			}

			projectName = words.ToSnakeCase(projectName)

			a.mkdir(projectName)
			a.mkdir(fmt.Sprintf("%s/app", projectName))
			a.mkdir(fmt.Sprintf("%s/app/routes", projectName))
			a.mkdir(fmt.Sprintf("%s/app/internal", projectName))

			if err = f.CreateReadme(projectName); err != nil {
				a.fatalError(err)
			}

			if err = f.CreateMakefile(projectName); err != nil {
				a.fatalError(err)
			}

			if err = f.CreateConfigFile(projectName); err != nil {
				a.fatalError(err)
			}

			if err = f.CreateSettings(projectName); err != nil {
				a.fatalError(err)
			}

			if err = f.CreateModels(projectName); err != nil {
				a.fatalError(err)
			}

			if err = f.CreatePyproject(projectName); err != nil {
				a.fatalError(err)
			}

			// Change the working directory to the app directory
			err = os.Chdir(projectName)
			if err != nil {
				a.fatalError(err)
			}

			a.installFastApi(projectName)
			a.installSqlModel(projectName)

			err = resource.GenerateResource("user")
			if err != nil {
				a.fatalError(err)
			}
		},
	}

	return cmd
}

// loader displays a loading animation
func loader(done chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Print(".")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
