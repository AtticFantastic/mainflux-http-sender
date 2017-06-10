package main

import (
	"fmt"
	"runtime"
	"time"
)

// const char* compilerVersion() {
// #if defined(__clang__)
// 	return __VERSION__;
// #elif defined(__GNUC__) || defined(__GNUG__)
// 	return "gcc " __VERSION__;
// #else
// 	return "non-gcc, non-clang (or an unrecognized version)";
// #endif
// }
import "C"

// TimeFormat is the reference format for build.Time. Make sure it stays in sync
// with the string passed to the linker in the root Makefile.
const TimeFormat = "2006/01/02 15:04:05"

var (
	// These variables are initialized via the linker -X flag in the
	// top-level Makefile when compiling release binaries.
	tag         = "unknown" // Tag of this build (git describe)
	utcTime     string      // Build time in UTC (year/month/day hour:min:sec)
	rev         string      // SHA-1 of this build (git rev-parse)
	cgoCompiler = C.GoString(C.compilerVersion())
	platform    = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
	version     = "unknown" // Version of the release, to be read from VERSION file
)

type (
	Info struct {
		GoVersion   string
		Tag         string
		Time        string
		Revision    string
		CgoCompiler string
		Platform    string
		Version     string
	}
)

// Short returns a pretty printed build and version summary.
func (b Info) Short() string {
	return fmt.Sprintf("Mainflux HTTP Sender %s (%s, built %s, %s)", b.Tag, b.Platform, b.Time, b.GoVersion)
}

// Timestamp parses the utcTime string and returns the number of seconds since epoch.
func (b Info) Timestamp() (int64, error) {
	val, err := time.Parse(TimeFormat, b.Time)
	if err != nil {
		return 0, err
	}
	return val.Unix(), nil
}

// GetInfo returns an Info struct populated with the build information.
func GetInfo() Info {
	return Info{
		GoVersion:   runtime.Version(),
		Tag:         tag,
		Time:        utcTime,
		Revision:    rev,
		CgoCompiler: cgoCompiler,
		Platform:    platform,
		Version:     version,
	}
}
