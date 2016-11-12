# GO crud template 

## About
This code is written with very little experience with Golang. I did my best. CRs and PRs always welcome :)

 - go 1.7,
 - gorilla mux for routing,
 - sqlx for database access (mysql),
 - tests using testify library
 - pkg/errors for better error handling
 - dockerized environment

## Usage

### API

urls:
 - GET http://localhost:8080/rest/banks/
 - GET http://localhost:8080/rest/banks/1
 - POST http://localhost:8080/rest/banks/ { "name": "BankName" }
 - PUT http://localhost:8080/rest/banks/1 { "name": "BankName" }
 - DELETE http://localhost:8080/rest/banks/1
 - DELETE http://localhost:8080/rest/banks/

### Docker

If you want to setup docker environemnt just use `./scripts/docker-compose.yml` with ![docker-compose](https://docs.docker.com/compose/)
 
## Hints

 - My $GOPATH=$HOME/go
 - Project dir: /home/users/kgoralski/go/src/github.com/kgoralski/go-crud-template

---
Many thanks to [@wendigo](https://github.com/wendigo) for a Code Review!
