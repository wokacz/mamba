package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wokacz/mamba/internal/words"
	"os"
)

func (app *application) initCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
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
			}

			projectName = words.ToSnakeCase(projectName)

			err = os.Mkdir(projectName, 0755)
			if err != nil {
				fmt.Println(err)
				return
			}

			err = f.CreateMakefile(projectName)
			err = f.CreateConfigFile(projectName)

			if err != nil {
				return
			}
		},
	}

	return cmd
}
