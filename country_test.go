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

