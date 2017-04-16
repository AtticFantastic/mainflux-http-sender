/**
 * Copyright (c) 2017 Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"net/http"
	"os"

	"github.com/mainflux/mainflux-http-sender/api"
)

const (
	help string = `
Usage: mainflux-influxdb [options]
Options:
	-a, --host	Host address
	-p, --port	Port
	-n, --nats	NATS host
	-q, --nport	NATS port
	-h, --help	Prints this message end exits`
)

type (
	Opts struct {
		HTTPHost string
		HTTPPort string

		NatsHost string
		NatsPort string

		Help bool
	}
)

func main() {
	opts := Opts{}

	flag.StringVar(&opts.HTTPHost, "a", "localhost", "HTTP server address.")
	flag.StringVar(&opts.HTTPPort, "p", "7070", "HTTP server port.")
	flag.StringVar(&opts.NatsHost, "n", "localhost", "NATS broker address.")
	flag.StringVar(&opts.NatsPort, "q", "4222", "NATS broker port.")
	flag.BoolVar(&opts.Help, "h", false, "Show help.")
	flag.BoolVar(&opts.Help, "help", false, "Show help.")

	flag.Parse()

	if opts.Help {
		fmt.Printf("%s\n", help)
		os.Exit(0)
	}

	// NATS
	api.NatsInit(opts.NatsHost, opts.NatsPort)

	// Print banner
	color.Cyan(banner)

	// Serve HTTP
	httpHost := fmt.Sprintf("%s:%s", opts.HTTPHost, opts.HTTPPort)
	http.ListenAndServe(httpHost, api.HTTPServer())
}

var banner = `
┌┬┐┌─┐┬┌┐┌┌─┐┬  ┬ ┬─┐ ┬   ┬ ┬┌┬┐┌┬┐┌─┐
│││├─┤││││├┤ │  │ │┌┴┬┘───├─┤ │  │ ├─┘
┴ ┴┴ ┴┴┘└┘└  ┴─┘└─┘┴ └─   ┴ ┴ ┴  ┴ ┴  
                                      
    == Industrial IoT System ==

    Made with <3 by Mainflux Team
[w] http://mainflux.io
[t] @mainflux

`
