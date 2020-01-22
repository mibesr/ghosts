package ghosts

import (
	"log"
	"os"
	"os/user"
	"path"
)

const CONFIG_HOME = ".config/ghosts"

func ConfigDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal("Cannot get current user: ", err)
	}
	confDir := path.Join(usr.HomeDir, CONFIG_HOME)
	if _, err := os.Stat(confDir); os.IsNotExist(err) {
		err = os.MkdirAll(confDir, 0700)
		if err != nil {
			log.Fatal("Cannot create configuration directory %s: %s", confDir, err)
		} else {
			log.Print("Configuration directory created")
		}
	}
	return confDir
}
