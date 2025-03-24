package files

func (f Files) Main() string {
	return `'''
Main module for the FastAPI application.
This module creates an instance of the FastAPI class and includes the routers defined in the routers module.
'''

from fastapi import FastAPI

from .internal.cors import configure_cors
from .internal.logs import configure_logging
from .routers import auth, users

app = FastAPI()
configure_logging()
configure_cors(app)

# Add the routers to the FastAPI app
app.include_router(auth.router)
app.include_router(users.router)
`
}
