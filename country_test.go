package gocountries

import (
  "testing"
  "log"
)

func TestCountriesByName(t *testing.T) {
  countries, err := CountriesByName("italy")
  if err != nil {
    t.Errorf("Got unexpected error for requested country 'italy': %v", err)
    return
  }

  if len(*countries) < 1 {
    t.Errorf("Unexpected: couldn't find any country with name 'italy'")
  }

  if (*countries)[0].Capital != "Rome" {
    t.Errorf("Unexpected capital for country 'italy'")
  }

  log.Printf("Fetched: %v", (*countries)[0])
}

func TestCountriesByCapital(t *testing.T) {
  capital := "tallinn"
  countries, err := CountriesByCapital(capital)
  if err != nil {
    t.Errorf("Got unexpected error for CountriesByCapital(\"%s\"): %v", capital, err)
    return
  }

  country := (*countries)[0]

  if country.Name != "Estonia" {
    t.Errorf("Got unexpected country: expected '%s', got '%s' instead", "Estonia", country.Name)
  }

  log.Printf("Fetched: %v", country)
}

