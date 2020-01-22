package main

import (
	"flag"
	"fmt"
)

func main() {
	convertPtr := flag.Bool("convert", false, "Convert from Java or NPM YML format to pmm")
	dbPtr := flag.String("sb", "", "a string")

	flag.Parse()

	fmt.Println("Will Convert:", *convertPtr)
	fmt.Println("Database:", *dbPtr)
}
