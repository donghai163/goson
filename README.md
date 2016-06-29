![logo](logo.svg "Goson")

The intuitive way to deal with JSON in Go. Inspired by SwiftyJSON.

# About
Goson is a very fast library.

## Install

```shell
go get github.com/panthesingh/goson
```

## Data

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

}

```

## Documentation

Documentation can be found on godoc:

https://godoc.org/github.com/panthesingh/goson

## Author

Panthe Singh, http://twitter.com/panthesingh
