# Gocountry
This is a [Go](golang.org) wrapper library around the API provided by [Restcountries](restcountries.eu).

## Installation
Just go with

`go get github.com/alediaferia/gocountries`

## Usage

```golang
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

## License
This library is provided with a MIT License.
