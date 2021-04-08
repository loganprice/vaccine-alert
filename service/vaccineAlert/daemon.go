package vaccineAlert

import (
	"log"
	"time"

	"github.com/loganprice/vaccine-alert/pkg/ifttt"
	"github.com/loganprice/vaccine-alert/pkg/types"
	"github.com/loganprice/vaccine-alert/pkg/vaccineSpotter"
)

func (s *Service) Daemon() error {
	var PerviousRun []types.AvailableSite
	ticker := time.NewTicker(time.Duration(s.Interval) * time.Second)
	for _ = range ticker.C {
		log.Println("Checking for vaccines")
		data, err := vaccineSpotter.GetData(s.State)
		if err != nil {
			return err
		}
		available, err := s.Parse(data)
		if err != nil {
			return err
		}
		newSites := Difference(available, PerviousRun)
		if len(newSites) == 0 {
			log.Println("No New Appointments Found")
		} else {
			log.Println("New Appointments found")
			err = ifttt.Notify(s.IFTTTEvent, s.IFTTTKey, newSites)
			if err != nil {
				return err
			}
		}
		PerviousRun = available
	}
	return nil
}
