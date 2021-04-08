package vaccineAlert

import (
	"github.com/loganprice/vaccine-alert/pkg/types"
	"github.com/loganprice/vaccine-alert/pkg/zip"
)

func (s *Service) Parse(site types.SiteInformation) ([]types.AvailableSite, error) {
	var Available []types.AvailableSite
	for _, location := range site.Features {
		if location.Properties.AppointmentsAvailable == true {
			if zip.Validate(location.Properties.PostalCode, s.AllowedZips) {
				VaccineTypes := s.GetVaccineTypes(location)
				if len(VaccineTypes) >= 1 {
					Available = append(Available, types.AvailableSite{
						ProviderBrand: location.Properties.ProviderBrand,
						Name:          location.Properties.Name,
						Address:       location.Properties.Address,
						City:          location.Properties.City,
						PostalCode:    location.Properties.PostalCode,
						VaccineBrand:  VaccineTypes,
						URL:           location.Properties.URL,
						LastFetched:   location.Properties.AppointmentsLastFetched,
					})
				}
			}
		}
	}
	return Available, nil
}

func (s *Service) GetVaccineTypes(location types.Feature) []string {
	var Brands []string
	if location.Properties.AppointmentVaccineTypes.JJ && s.JJ {
		Brands = append(Brands, "JJ")
	}
	if location.Properties.AppointmentVaccineTypes.Moderna && s.Moderna {
		Brands = append(Brands, "Moderna")
	}
	if location.Properties.AppointmentVaccineTypes.Pfizer && s.Pfizer {
		Brands = append(Brands, "Pfizer")
	}
	return Brands
}
