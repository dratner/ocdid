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

	results := d.ToStrings()

	log.Printf("Results: %+v", results)

	d = new(DivisionId)
	d.FromString("ocd-division/country:us/district:dc/school_district:youth_build_pcs_layc")
	log.Printf("OCDID: %+v", d)

	desc := d.ToDescription()
	log.Printf("Description: %s", desc)
}
