package files

import (
	"fmt"
	"os"
)

func (f Files) getSettingsContent(projectName string) string {
	return fmt.Sprintf(`import os

APP_NAME = '%s'
APP_VERSION = '0.1.0'

LOG_LEVEL = os.getenv('LOG_LEVEL', 'DEBUG')
LOG_FILE = os.getenv('LOG_FILE', 'app.log')
LOG_FORMAT = os.getenv('LOG_FORMAT', '%(asctime)s - %(name)s - %(levelname)s - %(message)s')
LOG_DATE_FORMAT = os.getenv('LOG_DATE_FORMAT', '%d-%m-%Y %H:%M:%S')

UVICORN_LOG_LEVEL = os.getenv('LOG_LEVEL', 'DEBUG')
FAST_API_LOG_LEVEL = os.getenv('LOG_LEVEL', 'DEBUG')
SQL_ALCHEMY_LOG_LEVEL = os.getenv('LOG_LEVEL', 'DEBUG')

API_PREFIX = '/api'
API_VERSION = 'v1'
API_TITLE = '%s'
API_DESCRIPTION = 'FastAPI application.'

CORS_ORIGINS = os.getenv('CORS_ORIGINS', '*')
CORS_METHODS = os.getenv('CORS_METHODS', 'GET, POST, PUT, DELETE')
CORS_HEADERS = os.getenv('CORS_HEADERS', 'Content-Type')

DATABASE_HOST = os.getenv('DATABASE_HOST', 'localhost')
DATABASE_PORT = os.getenv('DATABASE_PORT', '5432')
DATABASE_USER = os.getenv('DATABASE_USER', 'postgres')
DATABASE_PASSWORD = os.getenv('DATABASE_PASSWORD', 'password')
DATABASE_NAME = os.getenv('DATABASE_NAME', 'postgres')

JWT_SECRET = os.getenv('JWT_SECRET', 'secret')
JWT_ALGORITHM = os.getenv('JWT_ALGORITHM', 'HS256')
JWT_ACCESS_TOKEN_EXPIRATION = os.getenv('JWT_ACCESS_TOKEN_EXPIRATION', 3600)
JWT_REFRESH_TOKEN_EXPIRATION = os.getenv('JWT_REFRESH_TOKEN_EXPIRATION', 3600)`, projectName, projectName)
}

func (f Files) CreateSettings(projectName string) error {
	file, err := os.Create(fmt.Sprintf("%s/app/settings.py", projectName))
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

	content := f.getSettingsContent(projectName)

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}
