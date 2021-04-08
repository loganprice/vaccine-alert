package types

import "time"

type SiteInformation struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type     string `json:"type"`
	Geometry struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"geometry"`
	Properties struct {
		ID               int           `json:"id"`
		URL              string        `json:"url"`
		City             string        `json:"city"`
		Name             string        `json:"name"`
		State            string        `json:"state"`
		Address          string        `json:"address"`
		Provider         string        `json:"provider"`
		TimeZone         string        `json:"time_zone"`
		PostalCode       string        `json:"postal_code"`
		Appointments     []interface{} `json:"appointments"`
		ProviderBrand    string        `json:"provider_brand"`
		CarriesVaccine   bool          `json:"carries_vaccine"`
		AppointmentTypes struct {
		} `json:"appointment_types"`
		ProviderBrandID         int    `json:"provider_brand_id"`
		ProviderBrandName       string `json:"provider_brand_name"`
		ProviderLocationID      string `json:"provider_location_id"`
		AppointmentsAvailable   bool   `json:"appointments_available"`
		AppointmentVaccineTypes struct {
			JJ      bool `json:"jj"`
			Pfizer  bool `json:"pfizer"`
			Moderna bool `json:"moderna"`
		} `json:"appointment_vaccine_types"`
		AppointmentsLastFetched          time.Time `json:"appointments_last_fetched"`
		AppointmentsAvailableAllDoses    bool      `json:"appointments_available_all_doses"`
		AppointmentsAvailable2NdDoseOnly bool      `json:"appointments_available_2nd_dose_only"`
	} `json:"properties"`
}

type AvailableSite struct {
	ProviderBrand string
	VaccineBrand  []string
	Name          string
	Address       string
	City          string
	PostalCode    string
	URL           string
	LastFetched   time.Time
}
