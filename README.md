# pickme
Please follow the below instructions

1. install golang
2. execute follwing command on cmd:
  
  <br>//mysql library 
  <br>go get "github.com/go-sql-driver/mysql"
  
  <br>// routing library for go
  <br>go get "github.com/gin-gonic/gin"
  
  <br>//pickme serivce
  <br>go get "github.com/ashandk/pickme"
  
  <br>//moving in to the application
  <br>cd pickme

3. execute follwing command on cmd to migrate database
  <br>go run db_migrator.go
4. execute follwing command on cmd to run service
  <br>go run pickme_service.go



