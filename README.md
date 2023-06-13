# gosl – The Go Snippet Library

[![Go version][go_version_img]][go_dev_url]
[![Go report][go_report_img]][go_report_url]
![Code coverage][code_coverage_img]
[![License][license_img]][license_url]

The **Go Snippet Library** (_or **gosl** for a short_) provides **a snippet
collection** for working with routine operations in your **Go** programs with
a super **user-friendly** API and the most **efficient performance** (see
the [benchmarks][benchmarks] section).

## ⚡️ Quick start

Simply add `gosl` to your project:

```bash
go get github.com/koddr/gosl
```

Add the needed snippet to your Go program, like this:

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
can find on the [pkg.go.dev][go_dev_url] page.

The package provides two categories of functions: **regular** and **universal**
using generics (Go 1.18+). Also, note that some features will only work
correctly on Go 1.20 and above.

## 🔨 Regular functions

The regular functions of the `gosl` package are aimed at solving one single
task with the smallest possible allocation of your machine's resources.

### Concat

Concatenates strings `s` to the one string:

```go
s1 := "this "
s2 := "is "
s3 := "my string"

s := gosl.Concat(s1, s2, s3) // "this is my string"
```

### ContainsCaseInsensitive

Reports if string `substr` is within string `s` (case-insensitive by default):

```go
s := "Hello, WORLD!"
substr := "r"

b := gosl.ContainsCaseInsensitive(s, substr) // true
```

### IsFileExist

Reports whether a file exists on the specified `path`:

```go
p := filepath.Clean("~/Downloads/file.csv")

b := gosl.IsFileExist(p) // true|false
```

### IsDirExist

Reports whether a dir exists on the specified `path`:

```go
p := filepath.Clean("~/Downloads/my-folder")

b := gosl.IsDirExist(p) // true|false
```

### RandomString

Generates a (**really**) random string with a given size:

```go
size := 8

s, err := gosl.RandomString(size) // string, like "34f4ey7e"
if err != nil {
    log.Fatal(err)
}
```

### RenderStyled

Renders a styled string with a given `lipgloss.Style` template:

```go
tmpl := lipgloss.NewStyle().Foreground(lipgloss.Color("42")).Margin(1)

s := gosl.RenderStyled("This is a styled text", tmpl) // styled string
```

This function is a more comfortable wrapper for the
[charmbracelet/lipgloss][charmbracelet_lipgloss_url] library.

### ToString

Converts byte slice `b` to string or error:

```go
b := []byte("Hello, World!")

s, err := gosl.ToString(b) // "Hello, World!"
if err != nil {
    log.Fatal(err)
}
```

### ToBytes

Converts string `s` to byte slice or error:

```go
s := "Hello, World!"

b, err := gosl.ToBytes(s) // [48 65 6c ...]
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

### Equals

Compares two values of type `T`, return `true` if they are equal:

```go
s1 := "hello"
s2 := "hello"

b := gosl.Equals(s1, s2) // true
```

### NotEquals

Compares two values of type `T`, return `true` if they are **not** equal:

```go
s1 := 42
s2 := 64

b := gosl.NotEquals(s1, s2) // true
```

### ContainsInSlice

Reports if value `v` is within slice `s`:

```go
s := []string{"one", "two", "three"}
v := "two"

b := gosl.ContainsInSlice(s, v) // true
```

### ContainsInMap

Reports if key `k` is within map `m`:

```go
m := map[string]int{"one": 1, "two": 2, "three": 3}
k := "two"

b := gosl.ContainsInMap(m, k) // true
```

### ParseFileWithEnvToStruct

Parses the given file from `path` to struct `*T` with an (_optional_) environment variables for a secret data.

Set your secret data to environment variables with personal prefix (for ex., `MY_CONFIG`):

```console
export MY_CONFIG_TOKEN=my-secret-1234567
```

Create structured file in any of the supported file formats (JSON, YAML, TOML, or HCL) with the main data to parse (for
ex., `./config.yml`):

```yaml
url: https://my-server.com/api/v1
auth_type: Bearer
token: '{{ MY_CONFIG_TOKEN }}'
```

Create a new struct for a parsing data (for ex., `config`):

```go
type config struct {
    URL      string `koanf:"url"`
    AuthType string `koanf:"auth_type"`
    Token    string `koanf:"token"`
}
```

Add to your Go program:

```go
pathToFile := "./config.yml" // or any URL to file in the supported format
envPrefix := "MY_CONFIG"     // or "", if you don't want to use env
modelToParse := &config{}

cfg, err := gosl.ParseFileWithEnvToStruct(pathToFile, envPrefix, modelToParse)
if err != nil {
    log.Fatal(err)
}

// Results:
//  cfg.URL = "https://my-server.com/api/v1"
//  cfg.AuthType = "Bearer"
//  cfg.Token = "my-secret-1234567"
```

This generic function is based on the [knadh/koanf][knadh_koanf_url] library.

### Marshal

Marshal struct `user` to JSON data `j` (byte slice) or error:

```go
type user struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

u := &user{}

j, err := gosl.Marshal(u) // {"id": 0, "name": ""}
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

u, err := gosl.Unmarshal(j, m) // [id:1 name:Viktor]
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
BenchmarkEquals-8                               	319768486	         3.591 ns/op	       0 B/op	       0 allocs/op

BenchmarkNotEquals-8                            	1000000000	         0.5136 ns/op	       0 B/op	       0 allocs/op

BenchmarkConcat_String2-8                       	59083364	        19.91 ns/op	      32 B/op	       1 allocs/op
BenchmarkConcat_String8-8                       	27004447	        44.21 ns/op	     128 B/op	       1 allocs/op
BenchmarkConcat_String32-8                      	 9373778	       127.4 ns/op	     448 B/op	       1 allocs/op

BenchmarkToString_HelloWorld-8                  	100000000	        10.56 ns/op	      16 B/op	       1 allocs/op

BenchmarkToBytes_HelloWorld-8                   	1000000000	         0.6288 ns/op	   0 B/op	       0 allocs/op

BenchmarkRandomString_Size1-8                   	 3649489	       328.4 ns/op	       6 B/op	       3 allocs/op
BenchmarkRandomString_Size8-8                   	 3397297	       351.8 ns/op	      24 B/op	       3 allocs/op
BenchmarkRandomString_Size64-8                  	 2313856	       517.9 ns/op	     160 B/op	       3 allocs/op
BenchmarkRandomString_Size512-8                 	 1425562	       837.8 ns/op	    1280 B/op	       3 allocs/op
BenchmarkRandomString_Size4096-8                	  186254	      6331 ns/op	   10240 B/op	       3 allocs/op

BenchmarkMarshal_StructField_4-8                	 8584442	       139.9 ns/op	      48 B/op	       3 allocs/op
BenchmarkMarshal_StructField_16-8               	 2879486	       416.6 ns/op	     192 B/op	       3 allocs/op

BenchmarkUnmarshal_StructField_4-8              	 6960462	       169.3 ns/op	      32 B/op	       3 allocs/op
BenchmarkUnmarshal_StructField_16-8             	  774032	      1534 ns/op	     864 B/op	      45 allocs/op

BenchmarkRenderStyled-8                         	 1459971	       821.5 ns/op	     440 B/op	      12 allocs/op

BenchmarkContainsCaseInsensitive_HelloWorld-8   	24856041	        48.46 ns/op	      16 B/op	       1 allocs/op
BenchmarkContainsCaseInsensitive_LoremIpsum-8   	 1827114	       656.4 ns/op	     448 B/op	       1 allocs/op

BenchmarkContainsInSlice-8                      	122999034	         9.758 ns/op	   0 B/op	       0 allocs/op

BenchmarkContainsInMap-8                        	19123504	        62.61 ns/op	       0 B/op	       0 allocs/op

BenchmarkIsFileExist-8                          	  395916	      2941 ns/op	     240 B/op	       2 allocs/op

BenchmarkIsDirExist-8                           	  437505	      2696 ns/op	     224 B/op	       2 allocs/op
```

## 💡 Motivation

As you already know from my previous projects, I take an approach to software
development that makes **the developer's life totally easy**.

Why repeat the same snippet, for example, to translate a byte slice to a
string if you can add the most efficient solution to the library once and
import it where you need it? Exactly right! It's an unnecessary cognitive
load for those who will read your code in the future (and for you as well).

It is for these reasons that **The Go Snippet Library** (or `gosl` for a short)
provides a snippet collection for working with routine operations in your Go
programs with a super user-friendly API and the most efficient performance.

## 🏆 A win-win cooperation

And now, I invite you to participate in this project! Let's work **together** to
create the **largest** and **most useful** library of snippets for Go
programs on the web today.

- [Issues][repo_issues_url]: ask questions and submit your features.
- [Pull requests][repo_pull_request_url]: send your snippets or improvements
  to the current.

Your PRs & issues are welcome! Thank you 😘

## ⚠️ License

[`gosl`][repo_url] is free and open-source software licensed under the
[Apache 2.0 License][license_url], created and supported with 🩵 for people
and robots by [Vic Shóstak][author].

[go_version_img]: https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go
[go_report_img]: https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none
[go_report_url]: https://goreportcard.com/report/github.com/koddr/gosl
[go_dev_url]: https://pkg.go.dev/github.com/koddr/gosl
[code_coverage_img]: https://img.shields.io/badge/code_coverage-99%25-success?style=for-the-badge&logo=none
[license_img]: https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none
[license_url]: https://github.com/koddr/gosl/blob/main/LICENSE
[repo_url]: https://github.com/koddr/gosl
[repo_issues_url]: https://github.com/koddr/gosl/issues
[repo_pull_request_url]: https://github.com/koddr/gosl/pulls
[encoding_json_url]: https://pkg.go.dev/encoding/json

[charmbracelet_lipgloss_url]: https://github.com/charmbracelet/lipgloss

[knadh_koanf_url]: https://github.com/knadh/koanf

[benchmarks]: https://github.com/koddr/gosl/tree/main#%EF%B8%8F-benchmarks
[author]: https://github.com/koddr
