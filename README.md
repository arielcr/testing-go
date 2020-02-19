# Testing in Go

Sample code for testing in Go

### To run the test per package

```
go test -v --coverprofile cover.out github.com/arielcr/testing-go/api
```

```
go test -v --coverprofile cover.out github.com/arielcr/testing-go/math
```

### T get a report about the test

```
go tool cover -html=cover.out
```
