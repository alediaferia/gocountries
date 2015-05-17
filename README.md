# Gocountry
This is a [Go](https://golang.org) wrapper library around the API provided by [Restcountries](https://restcountries.eu).

## Installation
Just go with

`go get github.com/alediaferia/gocountries`

## Usage

```go
import (
  "github.com/alediaferia/gocountries"
  "fmt"
)

var countries *[]Country
countries, err = CountriesByName("italy")
if err == nil {
  c = (*countries)[0]
  fmt.Printf("The capital of Italy is %s", c.Capital)
}
```

## Contribution
Please, feel free to contribute to this project.

The following branches are currently active:

* ``master``: this is the stable branch which reflects the latest release
* ``develop``: this is where the magic happens :)

Additionally, I'd suggest you to create a hotfix/<word> branch
when contributing to this repo. Anyway, anything you like is fine,
just try to be as descriptive as possible when choosing your branch names
so that everything remains as readable as possible.

Thank you for your contributions in advance.

## License
This library is provided with a MIT License.
