# First: settin postgresql docker container

```
$ docker pull postgres
$ docker run --name go-sql-learn -e POSTGRES_PASSWORD=root -p 5433:5432 -d postgres

# go-sql-learn - container name; createdb - psql command for create DB; -U postgres - user of DB; gopg - name for new DB
$ docker exec -ti go-sql-learn createdb -U postgres gopg        
```