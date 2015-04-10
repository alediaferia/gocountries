# Gocountry
This is a [Go](golang.org) wrapper library around the API provided by [Restcountries](restcountries.eu).

## Installation
Just go with

`go get github.com/alediaferia/gocountries`

## Usage
```
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
````

## Contribution
Please, feel free to contribute to this project.

## License
This library is provided with a MIT License.
