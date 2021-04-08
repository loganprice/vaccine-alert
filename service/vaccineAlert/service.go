package vaccineAlert

import (
	"github.com/loganprice/vaccine-alert/pkg/types"
	"github.com/loganprice/vaccine-alert/pkg/zip"
)

type Service struct {
	Interval    int
	State       string
	AllowedZips types.ZipRadius
	IFTTTKey    string
	IFTTTEvent  string
	JJ          bool
	Pfizer      bool
	Moderna     bool
}

func New(zipKey, iftttEvent, iftttKey, zipCode string, interval, maxDistance int, jj, pfizer, moderna bool) (Service, error) {
	allowedZips, err := zip.GetRadius(zipKey, zipCode, maxDistance)
	if err != nil {
		return Service{}, err
	}
	zipInfo, err := zip.GetState(zipKey, zipCode)
	if err != nil {
		return Service{}, err
	}
	return Service{
		Interval:    interval,
		State:       zipInfo.State,
		AllowedZips: allowedZips,
		IFTTTEvent:  iftttEvent,
		IFTTTKey:    iftttKey,
		JJ:          jj,
		Pfizer:      pfizer,
		Moderna:     moderna,
	}, nil
}
