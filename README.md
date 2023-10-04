## Go Web Rankings

Go Web rankings is a simple web application responsible to rank websites. The main purpose of this project is to understand how to implement the Repository Pattern in Go.

### Requirements
* Docker
* Docker Compose
* Go 1.21+

### How To Run

1. Init the postgres container
```shell
make db
```

2. You can run the app using any of two postgres drivers (concrete dependencies). Choice one:
  * To run using the classic driver ("database/sql"):
      ```shell
      make run-classic
      ```
  * To run using the psx driver ("pgx/v4/pgxpool"):
      ```shell
      make run-pgx
      ```
  * To run using the gorm driver ("gorm.io/gorm"):
      ```shell
      make run-gorm
      ```
### Folder Structure


### Concepts to Understand
* "context" library
* "errors" library (specally *errors.Is* and *errors.As*)