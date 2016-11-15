This code is written with very little experience with Golang. I did my best. CRs and PRs always welcome :)

go-crud-template

 - go 1.7,
 - gorilla mux for routing,
 - sqlx for database access (mysql),
 - tests using testify library
 - pkg/errors for better error handling
 - spf13/viper for handling config files

urls:
 - GET http://localhost:8080/rest/banks/
 - GET http://localhost:8080/rest/banks/1
 - POST http://localhost:8080/rest/banks/ { "name": "BankName" }
 - PUT http://localhost:8080/rest/banks/1 { "name": "BankName" }
 - DELETE http://localhost:8080/rest/banks/1
 - DELETE http://localhost:8080/rest/banks/

Docker version 1.12.3:
 - Build  $docker build -t golang:go-app .
 - With mysqldb on host $docker run --name go-crud --network=host -it -d -p 8080:8080 golang:go-app
 - Should also work with mysql inside docker $docker run -d -p 3306:3306 --name mysql-db -e MYSQL_ROOT_PASSWORD=admin -d mysql:5.7
 
Hints:
 - My $GOPATH=$HOME/go
 - Project dir: /home/users/kgoralski/go/src/github.com/kgoralski/go-crud-template

---
Many thanks to [@wendigo](https://github.com/wendigo) for a Code Review!
