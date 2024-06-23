```
$ go get github.com/githubnemo/CompileDaemon
$ go get github.com/joho/godotenv

$ go get -u github.com/gin-gonic/gin

$ go get -u gorm.io/gorm
$ go get -u gorm.io/driver/postgres

# Start server
$ CompileDaemon -command="./gim-gorm-crud-api"
```


- DOCKER
docker run --name gin-crud -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

docker exec -it postgres createdb --username=root --owner=root simple_bank