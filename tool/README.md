

Install
===
- install on the source code directory
    
    ```go install```
    
- install on any path
    
    ```go install <repository-name><source-path>```
    
    example for $GO_PATH/src/github.com/user/hello.go
    
    ```go install github.com/user/hello```
    
Remote packages
===
- build and install the remote packages 
- if add to import, IDE tells how to add
```shell script
$ go get github.com/golang/example/hello
$ $GOPATH/bin/hello
```

Test
===
```shell script
go test
go test github.com/user/stringutil
```
