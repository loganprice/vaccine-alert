package ifttt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/loganprice/vaccine-alert/pkg/types"
)

func Notify(event, apiKey string, available []types.AvailableSite) error {
	for _, site := range available {
		bodyRaw := &types.IFTTT{
			Value1: fmt.Sprintf("%v - %v - %v", site.ProviderBrand, site.Name, site.VaccineBrand),
			Value2: site.URL,
			Value3: fmt.Sprintf("%v, %v, MN %v", site.Address, site.City, site.PostalCode),
		}
		byteBody, err := json.Marshal(bodyRaw)
		if err != nil {
			return err
		}
		http.Post(fmt.Sprintf("https://maker.ifttt.com/trigger/%v/with/key/%v", event, apiKey), "application/json", bytes.NewBuffer(byteBody))
	}
	return nil
}
