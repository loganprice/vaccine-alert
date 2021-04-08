package types

type ZipRadius struct {
	ZipCodes []struct {
		ZipCode  string  `json:"zip_code"`
		Distance float64 `json:"distance"`
		City     string  `json:"city"`
		State    string  `json:"state"`
	} `json:"zip_codes"`
}

type ZipInfo struct {
	ZipCode  string  `json:"zip_code"`
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
	City     string  `json:"city"`
	State    string  `json:"state"`
	Timezone struct {
		TimezoneIdentifier string `json:"timezone_identifier"`
		TimezoneAbbr       string `json:"timezone_abbr"`
		UtcOffsetSec       int    `json:"utc_offset_sec"`
		IsDst              string `json:"is_dst"`
	} `json:"timezone"`
	AcceptableCityNames []struct {
		City  string `json:"city"`
		State string `json:"state"`
	} `json:"acceptable_city_names"`
	AreaCodes []int `json:"area_codes"`
}
