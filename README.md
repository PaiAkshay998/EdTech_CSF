# EdTech_CSF

## Prerequisites
- Go 1.10 [Download link](https://golang.org/dl/#go1.10)
- MySQL

## Check prerequisites
- Check the go version installed.
```
go version
```
## Build instructions

- Download the repository and `cd` into it.
```
go get github.com/PaiAkshay998/Edtech_CSF
cd $GOPATH/src/github.com/PaiAkshay998/Edtech_CSF
```
- Create databases and run migrations
```
mysql -u root -p -e "CREATE DATABASE edtech_csf;"
migrate -url "mysql://root:YOUR_MYSQL_ROOT_PASSWORD@/edtech_csf" -path ./migrations up
```

- Fill in the database credentials in the `Dev` section of **config.json**.
- Run `go run main.go`

## Create Migrations
```
migrate -url "mysql://root:YOUR_MYSQL_ROOT_PASSWORD@/edtech_csf" -path ./migrations create migration_file_xyz
```


## Docker usage instructions
- Install [docker](https://docs.docker.com/engine/installation) and [docker-compose](https://docs.docker.com/compose/install).
- Fill in the *DB_NAME* and *DB_PASS* in *.env*. These are the credentials for the database container.
- Use the same credentials in `Docker` section *config.json* (*DbName* and *DbPassword*) and *docker-entry.sh* (in the `migrate` command).
- Run `docker-compose up`.
- Once the containers are up, you can get shell access by using
```
docker exec -it <CONTAINER_ID> bash
```
