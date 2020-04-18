# EdTech_CSF

Due to the national shutdown of schools to combat the COVID pandemic in India, there is immense potential to use technology to ensure students are learning at home. 

In an effort to classify and make EdTech solutions easily available to parents, teachers, and students to ensure continued teaching and learning at home, this website provides access to online resources for students and teachers so that learning can effectively take place at home. All solutions listed are free of cost and categorized for use by teachers and students/parents.

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

