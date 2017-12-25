package cfgfile

/*
 * this code runs both in parent and child
 * so beware of stdout availability (parent only)
 */

import (
	"os"
	"fmt"
	"errors"
	"github.com/go-ini/ini"
	. "github.com/aavzz/notifier/setup/cmdlnopts"
)

type beelineSection struct {
	login    string
	password string
	sender   string
}

type configurationFile struct {
	beeline beelineSection
}

var cfgFile1, cfgFile2 configurationFile
var cfgFile1ok, cfgFile2ok bool

// Package exported objects

func ReadConfig() {

	cfg, err := ini.Load(ConfigCfgFile())

	if err != nil {
		fmt.Printf("Cannot read configuration file %s: %s\n", ConfigCfgFile(), err)
		os.Exit(1)
	}

	cfgFile1ok=false
	_, err = cfg.GetSection("beeline")
	if err == nil {
		if cfg.Section("beeline").HasKey("login") {
			cfgFile1.beeline.login = cfg.Section("beeline").Key("login").String()
		}
		if cfg.Section("beeline").HasKey("password") {
			cfgFile1.beeline.password = cfg.Section("beeline").Key("password").String()
		}
		if cfg.Section("beeline").HasKey("sender") {
			cfgFile1.beeline.sender = cfg.Section("beeline").Key("sender").String()
		}
	}
	// true only means that we finished updating config data, not that the data is ok
	cfgFile1ok=true
	
	// update second copy of config data
	cfgFile2ok=false
	if err == nil {
		if cfg.Section("beeline").HasKey("login") {
			cfgFile1.beeline.login = cfg.Section("beeline").Key("login").String()
		}
		if cfg.Section("beeline").HasKey("password") {
			cfgFile1.beeline.password = cfg.Section("beeline").Key("password").String()
		}
		if cfg.Section("beeline").HasKey("sender") {
			cfgFile1.beeline.sender = cfg.Section("beeline").Key("sender").String()
		}
	}
	cfgFile2ok=true
}

func CfgFileContent() (*configurationFile, error) {
	if cfgFile1ok == true {
		c := &configurationFile{
			beeline: beelineSection{
				login: cfgFile1.beeline.login
				password: cfgFile1.beeline.password
				sender: cfgFile1.beeline.sender
			}
		}
		return c, nil
	}
	if cfgFile2ok == true {
		c := &configurationFile{
			beeline: beelineSection{
				login: cfgFile2.beeline.login
				password: cfgFile2.beeline.password
				sender: cfgFile2.beeline.sender
			}
		}
		return c, nil
	}
	return nil, errors.New("Error retrieving configuration file content")
}
