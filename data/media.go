package data

type MediaMeta struct {
	Files []struct {
		Sha256         string   `yaml:"sha256"`
		FileExtensions []string `yaml:"extensions"`
		Paths          []string `yaml:"extensions"`
		EarliestDate   int64    `yaml:"earliestDate"`
		ReviewComplete bool     `yaml:"reviewDone"`
		Ignore         bool     `yaml:"ignore"`
	}
}
