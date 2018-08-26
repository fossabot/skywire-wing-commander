// Copyright © 2018 BigOokie
//
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"path/filepath"

	"github.com/BigOokie/skywire-wing-commander/internal/utils"
	"github.com/BigOokie/skywire-wing-commander/internal/wcconfig"
	log "github.com/sirupsen/logrus"
)

type cmdlineFlags struct {
	dumpconfig bool
	version    bool
	help       bool
	about      bool
}

type wcBotApp struct {
	config   wcconfig.Config
	cmdFlags cmdlineFlags
}

// loadConfig manages the configuration load specifics
// offloading the detail from the `main()` funct
func (ba *wcBotApp) loadConfig() (c wcconfig.Config) {
	log.Debugln("loadConfig: Start")
	defer log.Debugln("loadConfig: Complete")
	// Load configuration
	c, err := wcconfig.LoadConfigParameters("config", filepath.Join(utils.UserHome(), ".wingcommander"), map[string]interface{}{
		"telegram.debug":                 false,
		"monitor.intervalsec":            10,
		"monitor.heartbeatintmin":        120,
		"monitor.discoverymonitorintmin": 120,
		"skymanager.address":             "127.0.0.1:8000",
		"skymanager.discoveryaddress":    "discovery.skycoin.net:8001",
	})

	if err != nil {
		log.Fatalf("Error loading configuration: %s", err)
		return
	}
	return
}

func (cf *cmdlineFlags) parseCmdLineFlags() {
	flag.BoolVar(&cf.version, "v", false, "print current version")
	flag.BoolVar(&cf.dumpconfig, "config", false, "print current config")
	flag.BoolVar(&cf.help, "help", false, "print application help")
	flag.BoolVar(&cf.about, "about", false, "print application information")
	flag.Parse()
	/*
		// if version cmd line flag `-v` then print version info and exit
		if versionFlag {
			fmt.Println(wcconst.BotAppVersion)
			fmt.Println("")
			os.Exit(0)
		}

		// if help cmd line flag `-help` then print version info and exit
		if helpFlag {
			fmt.Println(wcconst.MsgCmdLineHelp)
			fmt.Println("")
			os.Exit(0)
		}

		// if about cmd line flag `-about` then print version info and exit
		if aboutFlag {
			fmt.Println(wcconst.MsgAbout)
			fmt.Println("")
			os.Exit(0)
		}
	*/
}

func (cf *cmdlineFlags) handleCmdLineFlags() {

}

func (ba *wcBotApp) initLogging() {
	// Setup Log Formatter
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.DebugLevel)
}