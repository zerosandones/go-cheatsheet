# Go Cheat Sheet

## Go Compiler Commands

### Run

The _run_ command will compile and run the program. By default it does not create an executable file.

```bash
go run <module-name>/<file-name>
```

### Build

The _build_ command will compile the program but not run it. this will generate an executable file.

```bash
go build <module-name>/<file-name>
```

### Get

The _get_ command will download the library for usage.

```bash
go get <import-path>
```

## Modules

### Creating a Module

```go
go mod init <module-name>
```

You can then run you program by using the following command

```go
go run <module-name>
```

or

```go
go run .
```

### go.mod File

The _go.mod_ file is used to track the dependencies of your project. It is create when you use the _go mod init_ command.

If you are wanting to use a module that has not yet been published to a remote repository, you can use the _replace_ directive to point to a local directory.

```go
replace <import-path> => <local-path>
```

## Operators

### :=

The _:=_ operator is used to declare and initialize a variable. the type of the value assigned to the string tells the compiler what type the variable should be.

```go
var <variable-name> = <value>
```

Can be used instead of

```go
var <variable-name> <type> = <value>
```

## Functions

To define a function in Go, use the following syntax:

```go
func <function-name>(<parameters>) <return-type> {
	<function-body>
}
```

**NOTE:** Function names that start with a capital letter are exported and can be used outside the package.

Unlike other languages Go can return multiple values from a function.

If you want to be able to return an error, you can use the _error_ type.

```go
func <function-name>(<parameters>) (<return-type>, error) {
	<function-body>
}
```

To call and store the returned values from a function, use the following syntax:

```go
result, err := <function-name>(<parameters>)
```
