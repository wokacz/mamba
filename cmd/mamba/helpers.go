package main

import (
	"os"
	"os/exec"
)

// fatalError prints an error message and exits the program
func (a *application) fatalError(err error) {
	a.printError(err.Error())
	os.Exit(1)
}

// printSuccess prints a success message
func (a *application) printSuccess(message string) {
	a.cobra.Println("✅ ", message)
}

// printError prints an error message
func (a *application) printError(message string) {
	a.cobra.Println("🟥 ", message)
}

func (a *application) mkdir(path string) {
	err := os.Mkdir(path, 0755)
	if err != nil {
		a.fatalError(err)
	}
}

func (a *application) checkPython() (bool, string) {
	cmd := exec.Command("python", "--version")
	output, err := cmd.Output()

	if err != nil {
		a.printError("Python is not installed on this system.")
		return false, ""
	} else {
		a.printSuccess("Python is installed on this system. " + string(output))
		return true, string(output)
	}
}

// checkPoetry checks if Poetry is installed on the system
func (a *application) checkPoetry() (bool, string) {
	cmd := exec.Command("poetry", "--version")
	output, err := cmd.Output()

	if err != nil {
		a.printError("Poetry is not installed on this system.")
		return false, ""
	} else {
		a.printSuccess("Poetry is installed on this system. " + string(output))
		return true, string(output)
	}
}

func (a *application) installFastApi(projectName string) {
	cmd := exec.Command("poetry", "add", "fastapi[standard]")
	output, err := cmd.CombinedOutput()
	if err != nil {
		a.printError("Failed to install FastAPI: " + string(output))
		a.fatalError(err)
	} else {
		a.printSuccess("Successfully installed FastAPI: " + string(output))
	}
}

func (a *application) installSqlModel(projectName string) {
	cmd := exec.Command("poetry", "add", "sqlmodel")
	output, err := cmd.CombinedOutput()
	if err != nil {
		a.printError("Failed to install SQLModel: " + string(output))
		a.fatalError(err)
	} else {
		a.printSuccess("Successfully installed SQLModel: " + string(output))
	}
}
