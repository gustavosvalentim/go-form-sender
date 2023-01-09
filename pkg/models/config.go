package models

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Configuration struct {
	TemplatesDir string
}

var confpath string
var conf *Configuration

func NewConfigurationFrom(path string) *Configuration {
	cfg := new(Configuration)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return cfg
	}

	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range strings.Split(string(b), "\n") {
		kvpair := strings.Split(row, "=")
		if len(kvpair) > 2 {
			log.Fatal(fmt.Errorf("key %s has name or value with invalid character =", kvpair[0]))
		}
		switch kvpair[0] {
		case "templates_dir":
			cfg.TemplatesDir = kvpair[1]
		default:
		}
	}

	return cfg
}

func SetConfiguration(name, val string) {
	switch name {
	case "templates_dir":
		conf.TemplatesDir = val
	default:
		log.Fatal(fmt.Errorf("config does not exists %s", name))
	}
}

func PersistConfiguration() {
	f, err := os.OpenFile(confpath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Truncate(0); err != nil {
		log.Fatal(err)
	}

	w := bufio.NewWriter(f)
	if _, err := fmt.Fprintf(w, "templates_dir=%s\n", conf.TemplatesDir); err != nil {
		log.Fatal(err)
	}
	if err := w.Flush(); err != nil {
		log.Fatal(err)
	}
}

func InitConfiguration(path string) {
	confpath = path
	conf = NewConfigurationFrom(path)
}

func GetConfiguration() *Configuration {
	return conf
}
