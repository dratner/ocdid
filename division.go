package ocdid

import (
	"strings"
)

type DivisionId struct {
	Country               string
	State                 string
	County                string
	City                  string
	CongressionalDistrict string
	StateSenateDistrict   string
	StateHouseDistrict    string
}

func (d *DivisionId) countryToString() string {
	if len(d.Country) == 0 {
		return "us"
	}

	if len(d.Country) > 0 {
		// Implement country to code lookup
	}

	if len(d.Country) == 2 {
		return strings.Replace(strings.ToLower(d.Country), " ", "_", -1)
	}

	return "us"
}

func (d *DivisionId) StateString() string {

	country := d.countryToString()

	s := "ocd-division/country:" + country

	stateClause := getStateClause(d.State)

	if stateClause == "" {
		return ""
	}

	s += "/" + stateClause

	return s

}

func (d *DivisionId) ToStrings() []string {

	ids := make([]string, 0)

	country := d.countryToString()

	s := "ocd-division/country:" + country

	ids = append(ids, s)

	stateClause := getStateClause(d.State)

	if stateClause == "" {
		return ids
	}

	s += "/" + stateClause

	ids = append(ids, s)

	if d.County != "" {
		county := s + "/county:" + strings.Replace(strings.ToLower(d.County), " ", "_", -1)
		ids = append(ids, county)
	}

	if d.City != "" {
		place := s + "/place:" + strings.Replace(strings.ToLower(d.City), " ", "_", -1)
		ids = append(ids, place)
	}

	if d.CongressionalDistrict != "" {
		cong := s + "/cd:" + d.CongressionalDistrict
		ids = append(ids, cong)
	}

	if d.StateHouseDistrict != "" {
		shd := s + "/sldl:" + d.StateHouseDistrict
		ids = append(ids, shd)
	}

	if d.StateSenateDistrict != "" {
		sld := s + "/sldu:" + d.StateSenateDistrict
		ids = append(ids, sld)
	}

	return ids
}

func (d *DivisionId) FromString() error {
	return nil
}

/*
	CountyCode            string `json:"vb.vf_county_code"`
	JudicialDistrict      string `json:"vb.vf_judicial_district"`
	SchoolDistrict        string `json:"vb.vf_school_district"`
	CongressionalDistrict string `json:"vb.vf_cd"`
	CountyCouncil         string `json:"vb.vf_county_council"`
	StateSenateDistrict   string `json:"vb.vf_sd"`
	CityCouncil           string `json:"vb.vf_city_council"`
	PrecinctName          string `json:"vb.vf_precinct_name"`
	Township              string `json:"vb.vf_township"`
	CountyName            string `json:"vb.vf_county_name"`
	PrecinctId            string `json:"vb.vf_precinct_id"`
	StateHouseDistrict    string `json:"vb.vf_hd"`
	Ward                  string `json:"vb.vf_ward"`
	State                 string `json:"vb.vf_reg_cass_state"`
	Zip5                  string `json:"vb.vf_reg_cass_zip"`
	Zip4                  string `json:"vb.reg_cass_zip4"`
	MunicipalDistrict     string `json:"vb.vf_municipal_district"`
*/

/*

func (o *OCDId) Init() {
	if o.OCDType == "" {
		o.OCDType = "division"
	}

	if o.Country == "" {
		o.Country = "us"
	}
}

func (o *OCDId) FromString(s string) error {

	_ = strings.ToLower(s)

	return nil
}

func (o *OCDId) compose() string {

	s := fmt.Sprintf("ocd-%s/country:%s", o.OCDType, o.Country)

	o.State = strings.ToLower(o.State)

	if o.State == "dc" {
		s += "/district:dc"
	} else {
		s += "/state:" + o.State
	}

	if o.CongressionalDistrict != "" {
		s += "/cd:" + o.CongressionalDistrict
		return s
	}

	if o.StateSenateDistrict != "" {
		s += "/shdl:" + o.StateSenateDistrict
		return s
	}
	if o.StateHouseDistrict != "" {
		s += "/sldl:" + o.StateHouseDistrict
		return s
	}

	if o.County != "" {
		s += "/county:" + o.County
	} else if o.City != "" {
		s += "/place:" + o.City
	}

	return s
}

func (o *OCDId) ToString() string {

	var s string

	o.Init()
	s = o.compose()

	s = strings.Replace(strings.ToLower(o.compose()), " ", "_", -1)

	return s
}

}
*/
