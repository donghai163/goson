![Goson](https://dl.dropboxusercontent.com/u/9534337/goson_logo.svg "Goson")

A simple way to handle JSON data in Go.

### About
Goson was created to simplify reading JSON data within Go. Chain commands together
to traverse JSON structures with ease. This library has been inspired by [SwiftyJSON](https://github.com/SwiftyJSON/SwiftyJSON)

### Install

```shell
go get github.com/panthesingh/goson
```

### Starting

Create goson object from data. Returns an error if the data is not valid JSON.
```go
g, err := goson.Parse(data)
```
### Data

Every `Get()` call will return another Goson object. You can then access the data with a value function.
The default value types are `float64`, `int`, `bool`, `string`. If the key does not exist it the value function
will return the default zero value. To check if a key exists read the section on existence.

```go
name := g.Get("name").String()
age := g.Get("age").Int()
long := g.Get("longitude").Float()
booled := g.Get("isFat").Bool()

```
### Chaining
Chaining is very intuitive way to quickly traverse and manipulate the data you need.

```go
g.Get("key").Get("object").Index(0).Get("item").String()
```

### Existance
To check if a value exists use a type check on the `Value()` function. This returns
the Goson object's underlying `interface{}` value.

```go
if v, ok := g.Get("key").Value().(string); ok {
  println("the key value exists: ", v)
}
```

### Loop
Call `Len()` on any Goson object to length on the underlying value. Then you can loop through
the array using the `Index()` function.

```go
for i := 0; i < g.Len(); i++ {
    name := g.Index(i).Get("name").String()
    age := g.Index(i).Get("age").Int()
}
```

### Printing
Another useful feature within Goson is to print out the JSON structure at any value. Simply print any
Goson object to return the underlying structure in a pretty format.

```go
v := g.Get("sibling")
fmt.Println(v)
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

### Documentation

Documentation can be found on godoc:

https://godoc.org/github.com/panthesingh/goson

### Author

Panthe Singh, http://twitter.com/panthesingh
