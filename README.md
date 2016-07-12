![Goson](https://dl.dropboxusercontent.com/u/9534337/goson_logo.svg "Goson")

A simple and intuitive way to handle JSON in Go.

### About
Goson was created to simplify reading JSON within Go. 

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

Default value types are string, bool, float64 and int. Calling there corrosponding functions on a goson object 
will return the value if it exists or the default value. They do not return without knowing the function.

```go
name := g.Get("name").String()
age := g.Get("age").Int()
long := g.Get("longitude").Float()
booled := g.Get("isFat").Bool()

```
### Chaining
Goson allows an intuitive way to grab data as you would expect from a normal json operation.
Reading nested values is easy. The get function always returns a goson object which you can then call
get again and again until you reach the desired object.

```go
g.Get("key").Get("object").Index(0).Get("item").String()
```

### Existance
If you want to always check if the certain value exists you can use the conventional
style on the function "Value()"

```go
if v, ok := g.Get("key").Value().(string); ok {
  println("the key value exists: ", v)
}
```

### Loop
Goson allows you to use Len() to get the length of a supposed array and then using
the Index() function you can retrieve the element at the desired index.

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
