# Golang Library Command Line Project

## About the project

This application provides two featured library appliation from command line. This app take two commands such as **list** and **search**.

**List:** This command prints all books in the system.
```bash
foo@bar:~$ go run main.go list
1) Book1
2) Book2
foo@bar:~$ _
```
**Search:** This command is used with at least one another argument that is egual or longer tahn 3 char. It retuns books that contain string that given by you from console.
```bash
foo@bar:~$ go run main.go search Harry Potter
1) Harry Potter 1
2) Harry Potter 2
foo@bar:~$ _
```

In the other situations, program will print error messages and terminated.


## Notes


* Every project **MUST** use `dep` for vendor management and **MUST** checkin `vendor` direcotry.
*  To run application you have to change import statment about for searchhelper pacakge. You have two option.
*  First one is that you can put the searchelper directory in src directory that is in `GOPATH` . In this case you don't need to update import statment.
```go
package main
import (
	"fmt"
	"os"
	"path/filepath"
	"searchhelper"
	"strings"
)
```

* Second option is that you can set project path as a `GOPATH` and you can update import statment like below. 
```go
package main
import (
	"fmt"
	"os"
	"path/filepath"
	"my_packages/searchhelper"
	"strings"
)
```
* Also if there is any problem accsesing package or GOPATH you can set again GOPATH and you can use this command below and go visit this [link](https://stackoverflow.com/questions/68693154/package-is-not-in-goroot).
```bash
foo@bar:~$ go env -w GO111MODULE=off
```