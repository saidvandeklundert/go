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

! Compile the code AND run it:
go run main.go
go run *.go

! Run all files in current directory (open terminal in your project and then):
go run . 
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


Packages:
- all files that belong to a package should be in the same directory
- all files in a folder should belong to the same package

There are 2 type of packages in Go:
- Executable packages ( executable, not importable)
- Library packages: (not executable, importable)

Executable packages should always contain a function called main.

Scopes:
- function defined in the main package are usable throughout
- imported packages are visbile in that file only

A statement instructs Go to do something. 
An expression produces a value.
An operator can combine expressions.

Importing:

Create library:

Basic data types:

`int`: integer literals without a fractional part( -2 0 88 )
`float64`: integer literals wit a fractional part (-1.7 0. 88.8)
`bool`: pre-declared constants (true, false)
`string`: utf-8 1-4 bytes 

Declarations:

Assignments:

Type conversion: