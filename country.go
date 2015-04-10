package gocountries

import (
  "errors"
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "log"
)

const baseURL = "https://restcountries.eu/rest/v1/%s"

type Country struct {
  Name           string
  Capital        string
  AltSpellings   []string
  Relevance      string
  Region         string
  Subregion      string
  Translations   map[string]string
  Population     int32
  LatLng         []float32
  Demonym        string
  Area           float32
  Gini           float32
  Timezones      []string
  Borders        []string
  NativeName     string
  CallingCodes   []string
  TopLevelDomain []string
  Alpha2Code     string
  Alpha3Code     string
  Currencies     []string
  Languages      []string
}

func CountriesByName(name string) (*[]Country, error) {
  url := fmt.Sprintf(baseURL, fmt.Sprintf("name/%s", name))
  log.Printf("GET %s", url)
  res, err := http.Get(url)
  if err != nil {
    return nil, err
  }
  defer res.Body.Close() // So we don't forget about it later
  if res.StatusCode != 200 {
    e := errors.New(fmt.Sprintf("Unexpected API status code %s", res.Status))
    return nil, e
  }
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return nil, err
  }

  var c []Country
  err = json.Unmarshal(body, &c)
  if err != nil {
    return nil, err
  }
  return &c, nil
}

