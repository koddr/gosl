# gosl – The Go Snippet Library

<img src="https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go" alt="go version" />&nbsp;
<img src="https://img.shields.io/badge/code_coverage-98%25-success?style=for-the-badge&logo=none" alt="code coverage" />&nbsp;
<a href="https://goreportcard.com/report/github.com/koddr/gosl" target="_blank"><img src="https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none" alt="go report" /></a>&nbsp;
<img src="https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none" alt="license" />

The **Go Snippet Library** (_or **gosl** for a short_) provides **snippets** 
collection for working with routine operations in your **Go** programs with 
a super **user-friendly** API and the most efficient performance (see the 
[benchmarks](https://github.com/koddr/gosl/tree/main#%EF%B8%8F-benchmarks) section).

## ⚡️ Quick start

Install `gosl` package:

```bash
go get github.com/koddr/gosl
```

Add needed snippet to your Go program, like this:

```go
package main

import "github.com/koddr/gosl"

type user struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func main() {
    b := []byte("Hello, World!")
    
    s, err := gosl.ToString(b) // convert byte slice to string
    
    // ...
    
    json := []byte(`{"id":1,"name":"Viktor"}`)
    model := &user{}

    u, err := gosl.Unmarshal(json, model) // unmarshal JSON data to struct

    // ...
}
```

...or like this to have access to snippets as embedded struct:

```go
package main

import "github.com/koddr/gosl"

type App struct {
    // ...
    
    utils    *gosl.Utility                  // add regular snippets
    genUtils *gosl.GenericUtility[any, any] // add generic snippets
}

func (a *App) handleSomething() error {
    // ...
    
    s, err := a.utils.ToString(b) // convert byte slice to string
    
    // ...
    
    u, err := a.genUtils.Unmarshal(json, model) // unmarshal JSON data to struct
    
    // ...
}
```

## ✨ Usage

Basic usage and full code examples of all functions of the `gosl` package you can find on the [pkg.go.dev](https://pkg.go.dev/github.com/koddr/gosl) page.

The package provides two categories of functions: **regular** and **universal** using generics (Go 1.18+). Also, note that some features will only work correctly on Go 1.20 and above.

## 🔨 Regular functions

The regular functions of the `gosl` package are aimed at solving one single 
task with the smallest possible allocation of your machine's resources.

### ToString

Convert byte slice `b` to string or error:

```go
b := []byte("Hello, World!")

s, err := gosl.ToString(b)
if err != nil {
    log.Fatal(err)
}
```

### ToBytes

Convert string `s` to byte slice or error:

```go
s := "Hello, World!"

b, err := gosl.ToBytes(s)
if err != nil {
    log.Fatal(err)
}
```

### ContainsCaseInsensitive

Report if string `substr` is within string `s` (case-insensitive by default):

```go
s := "Hello, World!"
substr := "o"

b := gosl.ContainsCaseInsensitive(s, substr)
```

## 🛠️ Universal functions

The universal (or _generic_) functions of the `gosl` package are aimed at solving one 
particular task with the smallest possible allocation of your machine's 
resources, but can be applied to a huge number of user types.

> 💡 Hint: enjoy the benefits of using Go 1.18+ generics today! Instead of 
> writing a regular function for each of your types, just use **one generic 
> function** from the list below.

### Marshal

Marshal struct `user` to JSON data `j` (byte slice) or error:

```go
type user struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

u := &user{}

j, err := gosl.Marshal(u)
if err != nil {
    log.Fatal(err)
}
```

This generic function is a 100% compatible drop-in replacement for the standard 
[encoding/json](https://pkg.go.dev/encoding/json) library.

### Unmarshal

Unmarshal JSON data `j` (byte slice) to struct `user` or error:

```go
type user struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

j := []byte(`{"id":1,"name":"Viktor"}`)
m := &user{}

u, err := gosl.Unmarshal(j, m)
if err != nil {
    log.Fatal(err)
}
```

This generic function is a 100% compatible drop-in replacement for the standard 
[encoding/json](https://pkg.go.dev/encoding/json) library.

### ContainsInSlice

Report if value `v` is within slice `s`:

```go
s := []string{"one", "two", "three"}
v := "two"

b := gosl.ContainsInSlice(s, v)
```

### ContainsInMap

Report if key `k` is within map `m`:

```go
m := map[string]int{"one": 1, "two": 2, "three": 3}
k := "two"

b := gosl.ContainsInMap(m, k)
```

## ⏱️ Benchmarks

Run benchmarks on your machine by following command:

```bash
go test -v ./... -bench . -run ^$ -benchmem
```

And this is my results for all functions on test stand (Apple Macbook 
Air M1, 16 Gb RAM, macOS 13.3.1):

```bash
BenchmarkToString_HelloWorld-8                  	109688613	        10.58 ns/op	      16 B/op	       1 allocs/op

BenchmarkToBytes_HelloWorld-8                   	1000000000	       0.6284 ns/op	       0 B/op	       0 allocs/op

BenchmarkRandomString_Size1-8                   	 3708800	       324.4 ns/op	       6 B/op	       3 allocs/op
BenchmarkRandomString_Size8-8                   	 3504384	       342.5 ns/op	      24 B/op	       3 allocs/op
BenchmarkRandomString_Size64-8                  	 2303899	       517.5 ns/op	     160 B/op	       3 allocs/op
BenchmarkRandomString_Size512-8                 	 1437823	       833.8 ns/op	    1280 B/op	       3 allocs/op
BenchmarkRandomString_Size4096-8                	  189336	      6255 ns/op	   10240 B/op	       3 allocs/op

BenchmarkMarshal_StructField_4-8                	 8475444	       141.2 ns/op	      48 B/op	       3 allocs/op
BenchmarkMarshal_StructField_16-8               	 2847366	       421.5 ns/op	     192 B/op	       3 allocs/op

BenchmarkUnmarshal_StructField_4-8              	 7032111	       169.9 ns/op	      32 B/op	       3 allocs/op
BenchmarkUnmarshal_StructField_16-8             	  773876	      1553 ns/op	     864 B/op	      45 allocs/op

BenchmarkContainsCaseInsensitive_HelloWorld-8   	24499107	        48.53 ns/op	      16 B/op	       1 allocs/op
BenchmarkContainsCaseInsensitive_LoremIpsum-8   	 1804263	       663.4 ns/op	     448 B/op	       1 allocs/op

BenchmarkContainsInSlice-8                      	122393043	        9.817 ns/op	       0 B/op	       0 allocs/op

BenchmarkContainsInMap-8                        	19090268	        62.58 ns/op	       0 B/op	       0 allocs/op
```

## ⚠️ License

[`gosl`](https://github.com/koddr/gosl) is free and open-source software 
licensed under the [Apache 2.0 License](LICENSE), created and supported by
[Vic Shóstak](https://github.com/koddr).
