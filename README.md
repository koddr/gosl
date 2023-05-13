# gosl – The Go Snippet Library

![Go version][go_version_img]
[![Go report][go_report_img]][go_report_url]
![Code coverage][code_coverage_img]
[![License][license_img]][license_url]

The **Go Snippet Library** (_or **gosl** for a short_) provides **a snippet
collection** for working with routine operations in your **Go** programs with
a super **user-friendly** API and the most **efficient performance** (see
the [benchmarks][benchmarks] section).

## ⚡️ Quick start

Install `gosl` package:

```bash
go get github.com/koddr/gosl
```

Add needed snippet to your Go program, like this:

```go
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
import "github.com/koddr/gosl"

type App struct {
    // ...
    
    utils    *gosl.Utility                         // add regular snippets
    genUtils *gosl.GenericUtility[any, comparable] // add generic snippets
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

Basic usage and full code examples of all functions of the `gosl` package, you
can find on the [pkg.go.dev][gosl_go_dev_url] page.

The package provides two categories of functions: **regular** and **universal**
using generics (Go 1.18+). Also, note that some features will only work
correctly on Go 1.20 and above.

## 🔨 Regular functions

The regular functions of the `gosl` package are aimed at solving one single
task with the smallest possible allocation of your machine's resources.

### Concat

Concatenate strings `s` to the one string:

```go
s1 := "this "
s2 := "is "
s3 := "my string"

s := gosl.Concat(s1, s2, s3)
```

### ContainsCaseInsensitive

Report if string `substr` is within string `s` (case-insensitive by default):

```go
s := "Hello, WORLD!"
substr := "r"

b := gosl.ContainsCaseInsensitive(s, substr)
```

### RandomString

Generates a (**really**) random string with a given size:

```go
size := 8

s, err := gosl.RandomString(size)
if err != nil {
    log.Fatal(err)
}
```

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

## 🛠️ Universal functions

The universal (or _generic_) functions of the `gosl` package are aimed at
solving one
particular task with the smallest possible allocation of your machine's
resources, but can be applied to a huge number of user types.

> 💡 Hint: enjoy the benefits of using Go 1.18+ generics today! Instead of
> writing a regular function for each of your types, just use **one generic
> function** from the list below.

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
[encoding/json][encoding_json_url] library.

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
[encoding/json][encoding_json_url] library.

## ⏱️ Benchmarks

Run benchmarks on your machine by following command:

```bash
go test -v ./... -bench . -run ^$ -benchmem
```

And this is my results for all functions on test stand (Apple Macbook 
Air M1, 16 Gb RAM, macOS 13.3.1):

```bash
BenchmarkConcat_String2-8                       	58663996	        20.06 ns/op	      32 B/op	       1 allocs/op
BenchmarkConcat_String8-8                       	26829356	        44.16 ns/op	     128 B/op	       1 allocs/op
BenchmarkConcat_String32-8                      	 9321133	       127.8 ns/op	     448 B/op	       1 allocs/op

BenchmarkToString_HelloWorld-8                  	100000000	        10.56 ns/op	      16 B/op	       1 allocs/op

BenchmarkToBytes_HelloWorld-8                   	1000000000	         0.6288 ns/op	   0 B/op	       0 allocs/op

BenchmarkRandomString_Size1-8                   	 3488678	       344.6 ns/op	       6 B/op	       3 allocs/op
BenchmarkRandomString_Size8-8                   	 3394548	       353.3 ns/op	      24 B/op	       3 allocs/op
BenchmarkRandomString_Size64-8                  	 2313856	       517.9 ns/op	     160 B/op	       3 allocs/op
BenchmarkRandomString_Size512-8                 	 1423572	       838.9 ns/op	    1280 B/op	       3 allocs/op
BenchmarkRandomString_Size4096-8                	  185337	      6350 ns/op	   10240 B/op	       3 allocs/op

BenchmarkMarshal_StructField_4-8                	 8584442	       139.9 ns/op	      48 B/op	       3 allocs/op
BenchmarkMarshal_StructField_16-8               	 2838062	       420.8 ns/op	     192 B/op	       3 allocs/op

BenchmarkUnmarshal_StructField_4-8              	 6960462	       169.3 ns/op	      32 B/op	       3 allocs/op
BenchmarkUnmarshal_StructField_16-8             	  764182	      1553 ns/op	     864 B/op	      45 allocs/op

BenchmarkContainsCaseInsensitive_HelloWorld-8   	24856041	        48.46 ns/op	      16 B/op	       1 allocs/op
BenchmarkContainsCaseInsensitive_LoremIpsum-8   	 1797150	       695.9 ns/op	     448 B/op	       1 allocs/op

BenchmarkContainsInSlice-8                      	122999034	         9.758 ns/op	   0 B/op	       0 allocs/op

BenchmarkContainsInMap-8                        	19123504	        62.61 ns/op	       0 B/op	       0 allocs/op
```

## ⚠️ License

[`gosl`][gosl_url] is free and open-source software licensed under the
[Apache 2.0 License][license_url], created and supported with 🩵 for people
and robots by [Vic Shóstak][author].

[go_version_img]: https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go
[go_report_img]: https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none
[go_report_url]: https://goreportcard.com/report/github.com/koddr/gosl
[code_coverage_img]: https://img.shields.io/badge/code_coverage-98%25-success?style=for-the-badge&logo=none
[license_img]: https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none
[license_url]: https://github.com/koddr/gosl/blob/main/LICENSE
[gosl_go_dev_url]: https://pkg.go.dev/github.com/koddr/gosl
[encoding_json_url]: https://pkg.go.dev/encoding/json
[benchmarks]: https://github.com/koddr/gosl/tree/main#%EF%B8%8F-benchmarks
[gosl_url]: https://github.com/koddr/gosl
[author]: https://github.com/koddr
