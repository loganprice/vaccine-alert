package vaccineSpotter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/loganprice/vaccine-alert/pkg/types"
)

func GetData(state string) (types.SiteInformation, error) {
	res, err := http.Get(fmt.Sprintf("https://www.vaccinespotter.org/api/v0/states/%v.json", state))
	if err != nil {
		return types.SiteInformation{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return types.SiteInformation{}, err
	}
	var si types.SiteInformation
	err = json.Unmarshal(body, &si)
	if err != nil {
		return types.SiteInformation{}, err
	}
	return si, nil
}
