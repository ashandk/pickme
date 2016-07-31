# pickme
Please follow the below instructions

1. install golang
2. execute follwing command on cmd:
  go get "github.com/go-sql-driver/mysql"
  go get "github.com/gin-gonic/gin"
  go get "github.com/ashandk/pickme"
  cd pickme
3. execute follwing command on cmd to migrate database
  go run db_migrator.go
4. execute follwing command on cmd to run service
  go run pickme_service.go

