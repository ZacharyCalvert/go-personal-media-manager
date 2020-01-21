package old

import (
	"fmt"
	"io/ioutil"
	"os"

	. "github.com/ZacharyCalvert/go-personal-media-manager/data"

	"gopkg.in/yaml.v2"
)

type OldMediaFormat struct {
	Sha256         string   `yaml:"sha256"`
	FileExtensions []string `yaml:"extensions"`
	Paths          []string `yaml:"paths"`
	EarliestDate   int64    `yaml:"earliestDate"`
	ReviewComplete bool     `yaml:"reviewDone"`
	Ignore         bool     `yaml:"ignore"`
}

func Convert(database string) MediaMeta {
	var result MediaMeta

	return result
}

func ReadOld(database string) map[string]OldMediaFormat {
	_, err := os.Stat(database)
	if os.IsNotExist(err) {
		fmt.Printf("%s file does not exist.", database)
		panic(err)
	}

	data, err := ioutil.ReadFile(database)
	check(err)

	var result map[string]OldMediaFormat
	err = yaml.Unmarshal([]byte(data), &result)
	check(err)
	return result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
