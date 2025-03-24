package files

import (
	"fmt"
	"os"
)

func (f Files) getReadmeContent() string {
	return ``
}

func (f Files) CreateReadme(projectName string) error {
	file, err := os.Create(fmt.Sprintf("%s/README.md", projectName))
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

	content := f.getReadmeContent()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}
