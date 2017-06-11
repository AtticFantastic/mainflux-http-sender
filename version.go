package main

import (
	"fmt"
	"runtime"
)

// TimeFormat is the reference format for build.Time. Make sure it stays in sync
// with the string passed to the linker in the root Makefile.
const TimeFormat = "2006/01/02 15:04:05"

var (
	// These variables are initialized via the linker -X flag in the
	// top-level Makefile when compiling release binaries.
	Tag       = "unknown" // Tag of this build (git describe)
	Time      string      // Build time in UTC (year/month/day hour:min:sec)
	Revision  string      // SHA-1 of this build (git rev-parse)
	Platform  = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
	Version   = "unknown" // Version of the release, to be read from VERSION file
	GoVersion = runtime.Version()
)
