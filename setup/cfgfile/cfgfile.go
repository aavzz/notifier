package cfgfile

/*
 * this code runs both in parent and child
 * so beware of stdout availability (parent only)
 */

import (
	"os"
	"errors"
	"github.com/go-ini/ini"
	. "github.com/aavzz/notifier/setup/syslog"
	. "github.com/aavzz/notifier/setup/pidfile"
	. "github.com/aavzz/notifier/setup/cmdlnopts"
)

type beelineSection struct {
	login    string
	password string
	sender   string
}

type BeelineSection struct {
	Login    string
	Password string
	Sender   string
}

type emailSection struct {
	sender   string
}

type EmailSection struct {
	Sender   string
}

type configurationFile struct {
	beeline beelineSection
	email emailSection
}

type CfgFile struct {
	Beeline BeelineSection
	Email EmailSection
}


var cfgFile1, cfgFile2 configurationFile
var cfgFile1ok, cfgFile2ok bool

// Package exported objects



func ReadConfig() {

	cfg, err := ini.Load(ConfigCfgFile())

	if err != nil {
		daemonState := os.Getenv("_NOTIFY_DAEMON_STATE")
		if daemonState == "" {
			SysLog.Err(err.Error())
			RemovePidFile()
		}
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
	_, err = cfg.GetSection("email")
	if err == nil {
		if cfg.Section("email").HasKey("sender") {
			cfgFile1.email.sender = cfg.Section("email").Key("sender").String()
		}
	}
	// true only means that we finished updating config data, not that the data is ok
	cfgFile1ok=true
	
	// update second copy of config data
	cfgFile2ok=false
	_, err = cfg.GetSection("beeline")
	if err == nil {
		if cfg.Section("beeline").HasKey("login") {
			cfgFile2.beeline.login = cfg.Section("beeline").Key("login").String()
		}
		if cfg.Section("beeline").HasKey("password") {
			cfgFile2.beeline.password = cfg.Section("beeline").Key("password").String()
		}
		if cfg.Section("beeline").HasKey("sender") {
			cfgFile2.beeline.sender = cfg.Section("beeline").Key("sender").String()
		}
	}
	_, err = cfg.GetSection("email")
	if err == nil {
		if cfg.Section("email").HasKey("sender") {
			cfgFile2.email.sender = cfg.Section("email").Key("sender").String()
		}
	}
	cfgFile2ok=true
}

func CfgFileContent() (*CfgFile, error) {
	if cfgFile1ok == true {
		c := &CfgFile{
			Beeline: BeelineSect{
				Login: cfgFile1.beeline.login,
				Password: cfgFile1.beeline.password,
				Sender: cfgFile1.beeline.sender,
			},
			Email: EmailSect{
				Sender: cfgFile1.email.sender,
			},
		}
		return c, nil
	}
	if cfgFile2ok == true {
		c := &CfgFile{
			Beeline: BeelineSect{
				Login: cfgFile2.beeline.login,
				Password: cfgFile2.beeline.password,
				Sender: cfgFile2.beeline.sender,
			},
			Email: EmailSect{
				Sender: cfgFile2.email.sender,
			},
		}
		return c, nil
	}
	return nil, errors.New("Error retrieving configuration file content")
}
