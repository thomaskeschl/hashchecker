hashchecker
===========

Go program that with various flags that help you generate and confirm hashes for files using various hashing standards.

Setup
----------
Ensure the GOPATH binary directory is on your path, then:

```
go get github.com/thomaskeschl/hashchecker
hashchecker -h
```

If you would like to build the program as hc.exe, instead, do the following:

```
go get -d github.com/thomaskeschl/hashchecker
go build -o <GOPATH binary directory>/hc.exe github.com/thomaskeschl/hashchecker
hc -h
```
