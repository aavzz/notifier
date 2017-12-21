package cfgfile

/*
 * this code runs both in parent and child
 * so beware of stdout availability (parent only)
 */

import (
	"os"
	"fmt"
	"github.com/go-ini/ini"
	. "github.com/aavzz/notifier/setup/cmdlnopts"
)

type beelineSection struct {
	login    string
	password string
	senser   string
}

type configurationFile struct {
	beeline beelineSection
}

var cfgFile configurationFile

// Package exported objects

func ReadConfig() {

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

func ConfigBeelineLogin() string {
	return cfgFile.beeline.login
}

func ConfigBeelinePassword() uint16 {
	return cfgFile.beeline.password
}

func ConfigBeelineSender() uint16 {
	return cfgFile.beeline.sender
}
