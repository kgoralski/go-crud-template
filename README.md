go-crud-template

 - go 1.7,
 - gorilla mux for Routing,
 - sqlx for database access,
 - tests using testify library

urls:
 - GET http://localhost:8080/rest/banks/
 - GET http://localhost:8080/rest/banks/1
 - POST http://localhost:8080/rest/banks/ { "name": "BankName" }
 - PUT http://localhost:8080/rest/banks/1 { "id": 1, name": "BankName" }
 - DELETE http://localhost:8080/rest/banks/1

Docker:
 - Build  $docker build -t golang:go-app .
 - With mysqldb on host $docker run --name go-crud --network=host -it -d -p 8080:8080 golang:go-app
 
 Hints:
 - My $GOPATH=$HOME/work
 - Project dir: /home/users/kgoralski/work/src/github.com/kgoralski/go-crud-template/main