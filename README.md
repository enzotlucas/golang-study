# golang-study

##### Table of Contents  
- [Description](#description)
- [Notes](#notes)
    - [Main commands](#main-commands)
    - [Pointers](#pointers)
    - [Coding](#coding)
        - [Defer](#defer)

---

## Description
Repository i'm using to study GoLang.

## Notes
Notes of my study.

### Main commands
> Note: Every command needs to receive the main module after typed, example: go build main.go

Command | Description |
--------|-------------|
go build | Compiles the project
go run | Compiles and run the project
go fmt | Formats all code in the each file
go install | Installs the packages
go get | Downloads the raw source code of someone else's package
go test | Run tests

### Pointers

Command | Description
--- | ---
var example *type | You declare a variable that will reference a pointer to a type
*pointer | You get the value of a pointer
&variable | You get the reference of the pointer of a variable 

### Coding
Some notes about the code

#### Defer
Defer is a key word that makes the command be executed last on the method. You can use it to close a channel for example.
```go
func processFleet(ctx context.Context, fleet []Truck) error {
	var wg sync.WaitGroup
	errorsChan := make(chan error, len(fleet))
	defer close(errorsChan)

    //more code...
}
```