package setup

/*
 * this code runs both in parent and child
 * so beware of stdout availability (parent only)
 */

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
)

type smtpSection struct {
	address string
	port    uint16
}

type beelineSection struct {
	url      string
	login    string
	password string
}

type configurationFile struct {
	smtp    smtpSection
	beeline beelineSection
}

var cfgFile configurationFile

func readConfig() {

	cfgFile.smtp.address = "localhost"
	cfgFile.smtp.port = 25
	cfgFile.beeline.url = "https://beeline.amega-inform.ru/sendsms/"

	cfg, err := ini.Load(cmdLnOpts.cfgfile)

	if err != nil {
		fmt.Printf("Cannot read configuration file %s: %s\n", cmdLnOpts.cfgfile, err)
		os.Exit(1)
	}

	_, err = cfg.GetSection("smtp")

	if err == nil {
		if cfg.Section("smtp").HasKey("address") {
			cfgFile.smtp.address = cfg.Section("smtp").Key("address").String()
		}
		if cfg.Section("smtp").HasKey("port") {
			p, err := cfg.Section("smtp").Key("port").Uint()
			if err == nil {
				cfgFile.smtp.port = uint16(p)
			}
		}
	}

}
