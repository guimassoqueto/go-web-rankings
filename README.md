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

2. Run the application
```shell
make run
```

### Folder Structure


### Concepts to Understand
* "context" library
* "errors" library (specally *errors.Is* and *errors.As*)