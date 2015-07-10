# Gocountry

This is a [Go](https://golang.org) wrapper library around the API provided by
[Restcountries](https://restcountries.eu).

## Installation

Just go with

```shell
go get github.com/alediaferia/gocountries
```

## Example Usage

```go
package main

import (
    "fmt"
    "github.com/alediaferia/gocountries"
)

func main() {
    countries, err := gocountries.CountriesByName("italy")

    if err == nil {
        country := (countries)[0]
        fmt.Println(fmt.Sprintf("The capital of Italy is %s", country.Capital))
    }
}

```

## Contribution

Please, feel free to contribute to this project.

The following branches are currently active:

* ``master``: this is the stable branch which reflects the latest release
* ``develop``: this is where the magic happens :)

Additionally, I'd suggest you to create a hotfix/<word> branch when contributing
to this repo. Anyway, anything you like is fine, just try to be as descriptive
as possible when choosing your branch names so that everything remains as
readable as possible.

Thank you for your contributions in advance.

## License

This library is provided with a MIT License.
