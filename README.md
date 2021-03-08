# go

Dipping my toes into go.

Go is a compiled language.

Running a Go program:
- go build: compiles program so you can run it yourself
- go run: compiles and runs program at the same time


```
! Compile the code and then run it:
go build .\00_start\main.go         
.\main.exe

! Comile the code AND run it:
go run main.go
```

Making it portable:
```
OS X:
GOOS=darwin GOARCH=386 go build

Windows (cmd /c "set GOOS=darwin GOARCH=386 && go build"):
GOOS=windows GOARCH=386 go build

Linux:
GOOS=linux GOARCH=arm GOARM=7 go build
```

[play](https://play.golang.org/)

GOPATH:
```
go env GOPATH
```


fmt