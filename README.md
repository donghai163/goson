![Goson](https://dl.dropboxusercontent.com/u/9534337/goson_logo.svg "Goson")

A simple and intuitive way to handle JSON in Go.

### About
Goson was created to simplify reading JSON within Go. Chain commands together
to traverse complex JSON structures with ease. The values returned will never be nil unless explicity checked.


### Install

```shell
go get github.com/panthesingh/goson
```

### Getting Started

Create goson object from data. Returns an error if the data is not valid JSON.
```go
g, err := goson.Parse(data)
```
### Data

Use Get() as you would normally with a subscript on a map. Default value types are string, bool, float64 and int.
These values will never return nil if the value does not exist, instead you will get the default zero value.

```go
name := g.Get("name").String()
age := g.Get("age").Int()
long := g.Get("longitude").Float()
booled := g.Get("isFat").Bool()

```
### Chaining
Chaining is very intuitive way to quickly grab the data you need.

```go
g.Get("key").Get("object").Index(0).Get("item").String()
```

### Existance
To check if a certain value exists use Go's conventional way on the Value() function.

```go
if v, ok := g.Get("key").Value().(string); ok {
  println("the key value exists: ", v)
}
```

### Loop
Currently you can call Len() on any Goson object to return len(). Allowing you to loop
through JSON arrays with the Index() function.

```go
for i := 0; i < g.Len(); i++ {
    name := g.Index(i).Get("name").String()
    age := g.Index(i).Get("age").Int()
}
```

### Example

```go
package main

import (
  "github.com/panthesingh/jason"
)

func main() {

  json := `{
    "name": "Bob",
    "age": 100,
    "cars": [
      "Ferrari",
      "Lamborghini"
    ],
    "details": {
      "weight": 100
    }
  }`

  g, _ := goson.Parse([]byte(json))
  name := g.Get("name").String()
  age := g.Get("age").Int()
  cars := g.Get("cars")
  carOne := car.Index(0).String()
  carTwo := car.Index(1).String()
  weight := g.Get("details").Get("weight").String()

}

```

## Documentation

Documentation can be found on godoc:

https://godoc.org/github.com/panthesingh/goson

## Author

Panthe Singh, http://twitter.com/panthesingh
