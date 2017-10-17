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

func (d *DivisionId) ToStrings() map[string]string {

	ids := make(map[string]string, 0)

	country := d.countryToString()

	s := "ocd-division/country:" + country

	ids["country"] = s

	stateClause := getStateClause(d.State)

	if stateClause == "" {
		return ids
	}

	s += "/" + stateClause

	ids["state"] = s

	if d.County != "" {
		ids["county"] = s + "/county:" + strings.Replace(strings.ToLower(d.County), " ", "_", -1)
	}

	if d.City != "" {
		ids["place"] = s + "/place:" + strings.Replace(strings.ToLower(d.City), " ", "_", -1)
	}

	if d.CongressionalDistrict != "" {
		ids["cd"] = s + "/cd:" + d.CongressionalDistrict
		ids["exact"] = ids["cd"]
	}

	if d.StateHouseDistrict != "" {
		ids["sldl"] = s + "/sldl:" + d.StateHouseDistrict
		ids["exact"] = ids["sldl"]
	}

	if d.StateSenateDistrict != "" {
		ids["sldu"] = s + "/sldu:" + d.StateSenateDistrict
		ids["exact"] = ids["sldu"]
	}

	return ids
}

func (d *DivisionId) FromString(id string) error {

	parts := strings.Split(id, "/")

	var ci int
	var key, val string

	for _, part := range parts {
		ci = strings.Index(part, ":")
		if ci > 0 {
			key, val = part[:ci], part[ci+1:]
			switch key {
			case "district":
				d.State = val
				break
			case "territory":
				d.State = val
				break
			case "state":
				d.State = val
				break
			case "country":
				d.Country = val
				break
			case "county":
				d.County = val
				break
			case "cd":
				d.CongressionalDistrict = val
				break
			case "sldl":
				d.StateHouseDistrict = val
				break
			case "sldu":
				d.StateSenateDistrict = val
				break
			}
		}
	}
	return nil
}

func (d *DivisionId) ToDescription() string {
	var desc string

	if d.Country != "" {
		desc = d.Country
	} else {
		desc = "us"
	}

	if d.State != "" {
		desc = getStateName(d.State) + ", " + desc
	}

	if d.City != "" {
		desc = d.City + ", " + desc
	} else if d.County != "" {
		desc = d.County + " county, " + desc
	}

	return desc
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
