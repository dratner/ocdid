package ocdid

import (
	"strings"
)

func getStateClause(state string) string {

	if len(state) == 0 {
		return ""
	}

	if len(state) > 2 {
		state = getStateAbbrev(state)
	}

	if len(state) == 2 {
		state = strings.ToLower(state)
		t := abbrevToType(state)
		return strings.ToLower(t + ":" + state)
	}

	return ""
}

func getStateAbbrev(name string) string {
	sm := getStateMap()
	name = strings.Title(name)
	return sm[name]
}

func getStateName(abbrev string) string {
	sm := getStateMap()
	abbrev = strings.ToUpper(abbrev)
	for name, ab := range sm {
		if abbrev == ab {
			return name
		}
	}
	return ""
}

func abbrevToType(abbrev string) string {
	abbrev = strings.ToUpper(abbrev)
	types := map[string]string{"AL": "state",
		"AK": "state",
		"AS": "territory",
		"AZ": "state",
		"AR": "state",
		"CA": "state",
		"CO": "state",
		"CT": "state",
		"DE": "state",
		"DC": "district",
		"FM": "territory",
		"FL": "state",
		"GA": "state",
		"GU": "territory",
		"HI": "state",
		"ID": "state",
		"IL": "state",
		"IN": "state",
		"IA": "state",
		"KS": "state",
		"KY": "state",
		"LA": "state",
		"ME": "state",
		"MH": "territory",
		"MD": "state",
		"MA": "state",
		"MI": "state",
		"MN": "state",
		"MS": "state",
		"MO": "state",
		"MT": "state",
		"NE": "state",
		"NV": "state",
		"NH": "state",
		"NJ": "state",
		"NM": "state",
		"NY": "state",
		"NC": "state",
		"ND": "state",
		"MP": "territory",
		"OH": "state",
		"OK": "state",
		"OR": "state",
		"PW": "state",
		"PA": "state",
		"PR": "territory",
		"RI": "state",
		"SC": "state",
		"SD": "state",
		"TN": "state",
		"TX": "state",
		"UT": "state",
		"VT": "state",
		"VI": "territory",
		"VA": "state",
		"WA": "state",
		"WV": "state",
		"WI": "state",
		"WY": "state"}

	return types[abbrev]
}
func getStateMap() map[string]string {

	states := map[string]string{"alabama": "AL",
		"Alaska":                         "AK",
		"American Samoa":                 "AS",
		"Arizona":                        "AZ",
		"Arkansas":                       "AR",
		"California":                     "CA",
		"Colorado":                       "CO",
		"Connecticut":                    "CT",
		"Delaware":                       "DE",
		"District Of Columbia":           "DC",
		"Federated States Of Micronesia": "FM",
		"Florida":                        "FL",
		"Georgia":                        "GA",
		"Guam":                           "GU",
		"Hawaii":                         "HI",
		"Idaho":                          "ID",
		"Illinois":                       "IL",
		"Indiana":                        "IN",
		"Iowa":                           "IA",
		"Kansas":                         "KS",
		"Kentucky":                       "KY",
		"Louisiana":                      "LA",
		"Maine":                          "ME",
		"Marshall Islands":               "MH",
		"Maryland":                       "MD",
		"Massachusetts":                  "MA",
		"Michigan":                       "MI",
		"Minnesota":                      "MN",
		"Mississippi":                    "MS",
		"Missouri":                       "MO",
		"Montana":                        "MT",
		"Nebraska":                       "NE",
		"Nevada":                         "NV",
		"New Hampshire":                  "NH",
		"New Jersey":                     "NJ",
		"New Mexico":                     "NM",
		"New York":                       "NY",
		"North Carolina":                 "NC",
		"North Dakota":                   "ND",
		"Northern Mariana Islands":       "MP",
		"Ohio":           "OH",
		"Oklahoma":       "OK",
		"Oregon":         "OR",
		"Palau":          "PW",
		"Pennsylvania":   "PA",
		"Puerto Rico":    "PR",
		"Rhode Island":   "RI",
		"South Carolina": "SC",
		"South Dakota":   "SD",
		"Tennessee":      "TN",
		"Texas":          "TX",
		"Utah":           "UT",
		"Vermont":        "VT",
		"Virgin Islands": "VI",
		"Virginia":       "VA",
		"Washington":     "WA",
		"West Virginia":  "WV",
		"Wisconsin":      "WI",
		"Wyoming":        "WY"}

	return states
}
