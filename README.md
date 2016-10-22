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

