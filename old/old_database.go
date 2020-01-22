package old

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ZacharyCalvert/go-personal-media-manager/data"

	"gopkg.in/yaml.v2"
)

// LegacyMediaFormat represents the legacy media meta data for YML
type LegacyMediaFormat struct {
	Sha256         string   `yaml:"sha256"`
	FileExtensions []string `yaml:"extensions"`
	Paths          []string `yaml:"paths"`
	EarliestDate   int64    `yaml:"earliestDate"`
	ReviewComplete bool     `yaml:"reviewDone"`
	Ignore         bool     `yaml:"ignore"`
}

// Convert takes a path to a managed database in legacy format and converts it in place to pmm YML format
func Convert(database string) data.MediaMeta {
	var result data.MediaMeta

	return result
}

// ReadOld takes a path to a managed database in legacy format and loads the media metadata
func ReadOld(database string) map[string]LegacyMediaFormat {
	_, err := os.Stat(database)
	if os.IsNotExist(err) {
		fmt.Printf("%s file does not exist.", database)
		panic(err)
	}

	data, err := ioutil.ReadFile(database)
	check(err)

	var result map[string]LegacyMediaFormat
	err = yaml.Unmarshal([]byte(data), &result)
	check(err)
	return result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
