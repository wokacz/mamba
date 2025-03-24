package files

import (
	"fmt"
	"os"
)

// getMakefileContent
func (f *Files) getMakefileContent() string {
	return `# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## dev: start development server
.PHONY: dev
dev:
	@echo "Starting development server..."
	poetry run fastapi dev app/main.py
`
}

// CreateMakefile creates a Makefile in the current directory
func (f *Files) CreateMakefile(projectName string) error {
	file, err := os.Create(fmt.Sprintf("%s/Makefile", projectName))
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	content := f.GetMakefileContent()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}
