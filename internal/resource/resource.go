package resource

import (
	"fmt"
	"os"
)

func GenerateResource(resourceName string) error {
	err := createModel(resourceName)
	return err
}

func createModel(resourceName string) error {
	file, err := os.OpenFile("app/models.py", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	content := getModelContent(resourceName)

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}
