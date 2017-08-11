package ocdid

import (
	"log"
	"testing"
)

func TestDivision(t *testing.T) {

	d := new(DivisionId)
	d.State = "illinois"
	d.City = "chicago"
	d.County = "cook"
	d.CongressionalDistrict = "1"
	d.StateSenateDistrict = "2"
	d.StateHouseDistrict = "3"

	state := d.StateString()

	log.Printf("State: %s", state)

	ids := d.ToStrings()

	for _, ocdid := range ids {
		log.Printf("Id: %s", ocdid)
	}
}
