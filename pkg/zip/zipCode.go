package zip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/loganprice/vaccine-alert/pkg/types"
)

func GetRadius(apiKey, zipCode string, distance int) (types.ZipRadius, error) {
	res, err := http.Get(fmt.Sprintf("https://www.zipcodeapi.com/rest/%v/radius.json/%v/%v/mile", apiKey, zipCode, distance))
	if err != nil {
		return types.ZipRadius{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return types.ZipRadius{}, err
	}
	var radius types.ZipRadius
	json.Unmarshal(body, &radius)
	return radius, nil
}

func GetState(apiKey, zipCode string) (types.ZipInfo, error) {
	res, err := http.Get(fmt.Sprintf("https://www.zipcodeapi.com/rest/%v/info.json/%v/degrees", apiKey, zipCode))
	if err != nil {
		return types.ZipInfo{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return types.ZipInfo{}, err
	}
	var info types.ZipInfo
	json.Unmarshal(body, &info)
	return info, nil
}

func Validate(in string, allowedZips types.ZipRadius) bool {
	for _, zip := range allowedZips.ZipCodes {
		if in == zip.ZipCode {
			return true
		}
	}
	return false
}
