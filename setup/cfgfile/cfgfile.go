/*
Package cfgfile parses configuration file for the project
and creates internal representation of the parsed file.

If the configuration file cannot be read, the whole process terminates.

Internal representation is meant to be updated on signal (e.g. SIGHUP)
and is exported read-only in a consistent way.
*/
package cfgfile

import (
	"errors"
	. "github.com/aavzz/notifier/setup/cmdlnopts"
	. "github.com/aavzz/notifier/setup/pidfile"
	. "github.com/aavzz/notifier/setup/syslog"
	"github.com/go-ini/ini"
	"os"
)

// beelineSection holds common information needed to pass messages via beeline API.
// This is a private version, information is copied here from
// the configuration file and is exported to its public counterpart.
type beelineSection struct {
	login    string
	password string
	sender   string
}

// BeelineSection holds common information needed to pass messages via beeline API.
// This is a public version, information is copied here from
// its private counterpart.
type BeelineSection struct {
	Login    string
	Password string
	Sender   string
}

// configurationFile represents configuration file.
// This is a private version, information is copied here from
// configuration file and is exported to its public counterpart.
type configurationFile struct {
	beeline beelineSection
}

// CfgFile represents configuration file.
// This is a public version, information is copied here from its
// its public counterpart.
type CfgFile struct {
	Beeline BeelineSection
}

// Private structures and guardian vars.
var cfgFile1, cfgFile2 configurationFile
var cfgFile1ok, cfgFile2ok bool

// ReadConfig reads configuration file on startup or signal and fills
// private structures with the information it finds. In case of failure
// it logs a message and stops the process. Logging must have been
// initialized when ReadConfig is called.
func ReadConfig() {

	cfg, err := ini.Load(ConfigCfgFile())

	if err != nil {
		SysLog.Err(err.Error())
		RemovePidFile()
		os.Exit(1)
	}

	cfgFile1ok = false
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
	cfgFile1ok = true

	// update second copy of config data
	cfgFile2ok = false
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
	cfgFile2ok = true
}

// CfgFileContent exports internal representation of the configuration file
// in a safe and consistent way.
func CfgFileContent() (*CfgFile, error) {
	if cfgFile1ok == true {
		c := &CfgFile{
			Beeline: BeelineSection{
				Login:    cfgFile1.beeline.login,
				Password: cfgFile1.beeline.password,
				Sender:   cfgFile1.beeline.sender,
			},
		}
		return c, nil
	}
	if cfgFile2ok == true {
		c := &CfgFile{
			Beeline: BeelineSection{
				Login:    cfgFile2.beeline.login,
				Password: cfgFile2.beeline.password,
				Sender:   cfgFile2.beeline.sender,
			},
		}
		return c, nil
	}
	return nil, errors.New("Error retrieving configuration file content")
}
