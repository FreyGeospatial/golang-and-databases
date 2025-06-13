# golang-and-databases
investigating database connections with golang


To install external packages, we first need to initialize a module, which includes external dependencies:
```shell
go mod init golang-and-databases
```
Then, we can install the package(s):
```shell
go get github.com/jackc/pgx/v5 # this includes functionality to interact with postgres
```

Then, reference the package with the regular `import` command in our `go` code.