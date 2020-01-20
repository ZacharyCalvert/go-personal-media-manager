package pmm

import (
    "fmt"
    "flag"
)

type MediaMeta struct {
    Files []struct {
        Sha256 string `yaml:"sha256"`
        FileExtensions []string `yaml:"extensions"`
        Paths []string `yaml:"extensions"`
        EarliestDate int64 `yaml:"earliestDate"`
        ReviewComplete bool `yaml:"reviewDone"`
        Ignore bool `yaml:"ignore"`
    }
}

func main() {
    convertPtr := flag.Bool("convert", false, "Convert from Java or NPM YML format to pmm")
    dbPtr := flag.String("sb", "", "a string")

    flag.Parse()

    fmt.Println("Will Convert:", *convertPtr)
    fmt.Println("Database:", *dbPtr)
}
