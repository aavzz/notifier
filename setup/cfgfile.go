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

type beelineSection struct {
	login    string
	password string
}

type configurationFile struct {
	beeline beelineSection
}

var cfgFile configurationFile

func readConfig() {

	cfg, err := ini.Load(cmdLnOpts.cfgfile)

	if err != nil {
		fmt.Printf("Cannot read configuration file %s: %s\n", cmdLnOpts.cfgfile, err)
		os.Exit(1)
	}

	_, err = cfg.GetSection("beeline")

	if err == nil {
		if cfg.Section("beeline").HasKey("login") {
			cfgFile.beeline.login = cfg.Section("beeline").Key("login").String()
		}
		if cfg.Section("beeline").HasKey("password") {
			cfgFile.beeline.login = cfg.Section("beeline").Key("password").String()
		}
	}

}

// Package exported objects

/*
func ConfigSmtpServerAddress() string {
	return cfgFile.smtp.address
}

func ConfigSmtpServerPort() uint16 {
	return cfgFile.smtp.port
}*/
