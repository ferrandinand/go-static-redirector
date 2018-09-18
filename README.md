# go-static-redirector
Go static redirector is a simple golang app which can be used to deploy static files and manage all redirections in a JSON file.

## Static files
Static files must be placed in /public directory

## Redirections
Must be configured in a JSON file and this URL must be defined in REDIRECTION_FILE env var.


## RUN
``` 
export REDIRECTION_FILE=routes.json
export APP_PORT=8080
./go-static-redirector
``` 

## Tests
``` 
go tests
``` 
