package files

import (
	"fmt"
	"os"
)

func (f Files) getModelsContent() string {
	return `"""
This module contains the data models for the application.
The data models are defined using Pydantic data classes.
Each data model represents a specific type of data that the application will handle.
"""

from datetime import datetime
from uuid import UUID, uuid4

from sqlmodel import SQLModel, Field


# Define the base data models for the application

class BaseModel(SQLModel):
	id: UUID | None = Field(default_factory=uuid4, primary_key=True)


class TimestampModel(SQLModel):
	created_at: datetime | None = Field(default_factory=datetime.now, nullable=False)
	updated_at: datetime | None = Field(default_factory=datetime.now, nullable=False)


# Define the data models for the application`
}

func (f Files) CreateModels(projectName string) error {
	file, err := os.Create(fmt.Sprintf("%s/app/models.py", projectName))
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

	content := f.getModelsContent()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}
