package gocountries

import (
  "errors"
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
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

func doRestcountriesCall(apiSuffix string) ([]byte, error) {
  url := fmt.Sprintf(baseURL, apiSuffix)
  res, err := http.Get(url)
  if err != nil {
    return []byte{}, err
  }
  defer res.Body.Close()
  if res.StatusCode != 200 {
    e := errors.New(fmt.Sprintf("Unexpected API status code %s", res.Status))
    return []byte{}, e
  }
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return []byte{}, err
  }
  return body, nil
}

func CountriesByName(name string) ([]Country, error) {
  resData, err := doRestcountriesCall(fmt.Sprintf("name/%s", name))

  if err != nil {
    return nil, err
  }
  var c []Country
  err = json.Unmarshal(resData, &c)
  if err != nil {
    return nil, err
  }
  return c, nil
}

func CountriesByCapital(name string) ([]Country, error) {
  resData, err := doRestcountriesCall(fmt.Sprintf("capital/%s", name))

  if err != nil {
    return nil, err
  }
  var c []Country
  err = json.Unmarshal(resData, &c)
  if err != nil {
    return nil, err
  }
  return c, nil
}

