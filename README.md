[![GoDoc](https://godoc.org/github.com/owulveryck/gohaystack?status.svg)](https://pkg.go.dev/github.com/owulveryck/gohaystack?tab=doc) [![Go Report Card](https://goreportcard.com/badge/github.com/owulveryck/gohaystack)](https://goreportcard.com/report/github.com/owulveryck/gohaystack)
![](https://github.com/owulveryck/gohaystack/workflows/Go/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/owulveryck/gohaystack/badge.svg?branch=master)](https://coveralls.io/github/owulveryck/gohaystack?branch=master)

# gohaystack

This is a toy library to play with [Project Haystack](https://project-haystack.org/) in Go. It only relies on Go's standard library.

## What is project-haystack?

As stated in the web site

> "`Project Haystack`_ is an open-source initiative to streamline
> working with data from the Internet of Things. We standardize
> semantic data models and web services with the goal of making
> it easier to unlock value from the vast quantity of data being
> generated by the smart devices that permeate our homes, buildings,
> factories, and cities. Applications include automation, control,
> energy, HVAC, lighting, and other environmental systems."
>
> -- Project-Haystack

## What's included

This library is rather incomplete, as I am using it only for testing purpose (discovering haystack, building a sample server, and transform data from a format to another)
This library has a `Grid` structure that can be (de)serialized to and from haystack format in JSON (no zinc yet).
The `Grid`  hold an array of `Entity`.
An `Entity` has an `id` (format `HaystackID` which is a pointer to a string) a `Dis` field (a string used for display purpose) and a hash map of tags.

a tag is a composition of a `*Label` and a `*Value`

Please, see the API documentation for more information.

| Haystack kind | JSON Serialize | JSON Deserialize | Zinc Serialize | Zinc Deserialize |
|---------------|----------------|------------------|----------------|------------------|
| Grid          | x              | x                | x              |                  |
| List          | x              | x                | x              |                  |
| Dict          | x              | x                | x              |                  |
| Null          | x              | x                | x              |                  |
| Bool          | x              | x                | x              |                  |
| Marker        | x              | x                | x              |                  |
| Remove        | x              | x                | x              |                  |
| NA            | x              | x                | x              |                  |
| Number        | x              | x                | x              |                  |
| Ref           | x              | x                | x              |                  |
| Str           | x              | x                | x              |                  |
| Date          | x              | x                | x              |                  |
| Time          | x              | x                | x              |                  |
| DateTime      | x              | x                | x              |                  |
| URI           | x              | x                | x              |                  |
| Coord         | x              | x                | x              |                  |
| Xstr          |                |                  |                |                  |

## Haystack tags package

the [`tags`](tags) subpackage contains the label and values definitions of the official haystack project tag list. This package is a commodity.

## Example

### Basic example

Generate a grid and add simple values:

```go
g := NewGrid()
myTagLabel := NewLabel("mytag")
myTagSite := NewLabel("site")
myTagLabel.Display = "the display"
mySite := NewHaystackID("myreference")
entity := g.NewEntity(mySite)
myTagValue := NewStr("foo")
entity.SetTag(myTagLabel, myTagValue)
entity.SetTag(myTagSite, MarkerValue)
```

### JSON

Generate a grid from a json payload:
```go
g := grid.NewGrid()
dec := json.NewDecoder(os.Stdin)
err := dec.Decode(&g)
if err != nil {
    log.Fatal(err)
}
```

## Complete example

the `example/graph` contains a CLI tool that reads a JSON encoded haystack grid from stdin and generates a dot compatible representation on stdout.

## Caution

This is a toy project given without any warranty. I made it for my own need to discover haystack.
If you want to take over the maintenance, feel free to contact me.
