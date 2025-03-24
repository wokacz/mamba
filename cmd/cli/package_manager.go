package main

import "os/exec"

func (app *application) checkPoetry() (bool, string) {
	cmd := exec.Command("poetry", "--version")
	output, err := cmd.Output()

	if err != nil {
		app.printError("Poetry is not installed on this system.")
		return false, ""
	} else {
		app.printSuccess("Poetry is installed on this system." + string(output))
		return true, string(output)
	}
}
