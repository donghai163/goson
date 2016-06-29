![logo](https://us-bucket-host.s3.amazonaws.com/jason/jason_cropped_4.svg)
The better way to deal with JSON data in Go. Insipired by SwiftyJSON 

# About

Jason is designed to be convenient for reading arbitrary JSON while still honoring the strictness of the language. Inspired by other libraries and improved to work well for common use cases. It currently focuses on reading JSON data rather than creating it. [API Documentation](http://godoc.org/github.com/antonholmquist/jason) can be found on godoc.org.

## Install

```shell
go get github.com/panthesingh/goson
```

## Data types

The following golang values are used to represent JSON data types. It is consistent with how `encoding/json` uses primitive types.

- `bool`, for JSON booleans
- `json.Number/float64/int64`, for JSON numbers
- `string`, for JSON strings
- `[]*Value`, for JSON arrays
- `map[string]*Value`, for JSON objects
- `nil` for JSON null

## Examples

### Create from bytes

Create goson object from data. Returns an error if the data is not valid JSON.

```go
g, err := goson.Parse(data)

```

### Read values

Reading values is easy. If the key path is invalid or type doesn't match, it will return an error and the default value.

```go
name := g.Get("name").String()
age := g.Get("age").Int()
long := g.Get("longitude").Float()
booled := g.Get("isFat").Bool()

```

### Chaining

Reading nested values is easy. If the path is invalid or type doesn't match, it will return the default value and an error.

```go
g.Get("bob").Get("one").Index(0).Get("Money").String()

```

### Loop through array

Looping through an array is done with `GetValueArray()` or `GetObjectArray()`. It returns an error if the value at that keypath is null (or something else than an array).

```go
for i := 0; i < g.Len(); i++ {
    name := g.Index(i).Get("name").String()
    age := g.Index(i).Get("age").Int()
}
```

## Sample

Example:

```go
package main

import (
  "github.com/antonholmquist/jason"
  "log"
)

func main() {

  json := `{
    "name": "Walter White",
    "age": 51,
    "children": [
      "junior",
      "holly"
    ],
    "other": {
      "occupation": "chemist",
      "years": 23
    }
  }`

  g, err := goson.Parse([]byte(json))

  name := v.GetString("name")
  age := v.GetNumber("age")
  occupation, _ := v.GetString("other", "occupation")
  years, _ := v.GetNumber("other", "years")

  log.Println("age:", age)
  log.Println("name:", name)
  log.Println("occupation:", occupation)
  log.Println("years:", years)

  children, _ := v.GetStringArray("children")
  for i, child := range children {
    log.Printf("child %d: %s", i, child)
  }

  others, _ := v.GetObject("other")

  for _, value := range others.Map() {

    s, sErr := value.String()
    n, nErr := value.Number()

    if sErr == nil {
      log.Println("string value: ", s)
    } else if nErr == nil {
      log.Println("number value: ", n)
    }
  }
}

```

## Documentation

Documentation can be found on godoc:

https://godoc.org/github.com/panthesingh/goson

## Author

Panthe Singh, http://twitter.com/panthesingh
