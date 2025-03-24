package files

import (
	"fmt"
	"os"
)

func (f Files) getPyprojectContent(projectName string) string {
	return fmt.Sprintf(`[project]
name = "%s"
description = ""
version = "0.1.0"
package-mode = false
readme = "README.md"
authors = []
requires-python = ">=3.13"

[build-system]
requires = ["poetry-core>=2.0.0,<3.0.0"]
build-backend = "poetry.core.masonry.api"
`, projectName)
}

func (f Files) CreatePyproject(projectName string) error {
	file, err := os.Create(fmt.Sprintf("%s/pyproject.toml", projectName))
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

	content := f.getPyprojectContent(projectName)

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}
