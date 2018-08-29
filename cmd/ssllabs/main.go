// main.go
//
// Copyright 2018 © by Ollivier Robert <roberto@keltia.net>

/*
This is just a very short example.
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/keltia/ssllabs"
)

const (
	// MyVersion is for the app
	MyVersion = "0.1.0"
)

var (
	fDebug       bool
	fDetailed    bool
	fVerbose     bool
	fShowVersion bool

	// MyName is the application name
	MyName = filepath.Base(os.Args[0])
)

func init() {
	flag.BoolVar(&fDetailed, "d", false, "Get a detailed report")
	flag.BoolVar(&fVerbose, "v", false, "Verbose mode")
	flag.BoolVar(&fDebug, "D", false, "Debug mode")
	flag.BoolVar(&fShowVersion, "V,version", false, "Display version & exit.")
	flag.Parse()

	if fShowVersion {
		fmt.Fprintf(os.Stderr, "%s version %s API v3\n",
			MyName, ssllabs.Version())
		os.Exit(0)
	}

	if len(flag.Args()) == 0 {
		log.Fatalf("You must give at least one site name!")
	}
}

func main() {
	var level = 0

	site := flag.Arg(0)

	if fVerbose {
		level = 1
	}

	if fDebug {
		level = 2
		fVerbose = true
	}

	// Setup client
	c, err := ssllabs.NewClient(ssllabs.Config{Log: level})
	if err != nil {
		log.Fatalf("error setting up client: %v", err)
	}

	if fDetailed {

		report, err := c.GetDetailedReport(site)
		if err != nil {
			log.Fatalf("impossible to get grade for '%s'\n", site)
		}

		// Just dump the json
		fmt.Printf("%v\n", report)
	} else {
		fmt.Printf("%s Wrapper: %s API version %s\n\n",
			MyName, MyVersion, ssllabs.Version())
		grade, err := c.GetGrade(site)
		if err != nil {
			log.Fatalf("impossible to get grade for '%s': %v\n", site, err)
		}
		fmt.Printf("Grade for '%s' is %s\n", site, grade)
	}
}
