package cmd

import (
	"github.com/loganprice/vaccine-alert/pkg/types"
	"github.com/loganprice/vaccine-alert/service/vaccineAlert"

	"github.com/urfave/cli/v2"
)

var PerviousRun []types.AvailableSite

func Run(c *cli.Context) error {
	s, err := vaccineAlert.New(c.String("zip-key"), c.String("ifttt-event"), c.String("ifttt-key"), c.String("zip-code"), c.Int("interval"), c.Int("max-distance"), c.Bool("jj-allowed"), c.Bool("pfizer-allowed"), c.Bool("moderna-allowed"))
	if err != nil {
		return err
	}
	err = s.Daemon()
	if err != nil {
		return err
	}
	return nil
}
