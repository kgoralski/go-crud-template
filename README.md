# Golang CRUD Application Scaffold

## About
This code is written with very little experience with Golang. I did my best. CRs and PRs always welcome :)

 - go 1.7,
 - gorilla mux for routing,
 - sqlx for database access (mysql),
 - tests using testify library
 - pkg/errors for better error handling
 - spf13/viper for handling config files 
 - glide for package management
 - dockerized environment

## Usage

### API

 - GET http://localhost:8080/rest/banks/
 - GET http://localhost:8080/rest/banks/1
 - POST http://localhost:8080/rest/banks/ { "name": "BankName" }
 - PUT http://localhost:8080/rest/banks/1 { "name": "BankName" }
 - DELETE http://localhost:8080/rest/banks/1
 - DELETE http://localhost:8080/rest/banks/

### docker-compose

If you want to setup docker environemnt just use `./scripts/docker-compose.yml` with [docker-compose](https://docs.docker.com/compose/).

Go to `./scripts` directory and execute

```
# start docker environment
$ docker-compose up -d (--build)

# list running services
$ docker-compose ps

# stop all containers
$ docker-compose stop

# remove all containers
$ docker-compose rm
```
 
### single docker
 - Build app: &docker build -t golang:go-app .  && $docker run --name go-crud --network=host -it -d -p 8080:8080 golang:go-app
 - You can run localDB with: $docker run -d -p 3306:3306 --name mysql-db -e MYSQL_ROOT_PASSWORD=admin -d mysql:5.7
 
### Run Locally:
 - start mysql on 3306 port (and execute db script - manually currently)
 - start app from IDE or after install with flag -env=dev and giving -configFilePath, ./go-crud-template -env=dev
 
## Hints
 - My $GOPATH=$HOME/go
 - Project dir: /home/users/kgoralski/go/src/github.com/kgoralski/go-crud-template

---
Many thanks to [@wendigo](https://github.com/wendigo) for a Code Review!
