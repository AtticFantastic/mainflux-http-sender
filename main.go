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
	"log"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/fatih/color"

	"github.com/cenkalti/backoff"
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

		Help    bool
		Version bool
	}
)

var (
	opts Opts
)

func tryNatsConnect() error {
	var err error

	log.Print("Connecting to NATS... ")
	err = api.NatsInit(opts.NatsHost, opts.NatsPort)
	return err
}

func main() {
	flag.StringVar(&opts.HTTPHost, "a", "0.0.0.0", "HTTP server address.")
	flag.StringVar(&opts.HTTPPort, "p", "7070", "HTTP server port.")
	flag.StringVar(&opts.NatsHost, "n", "0.0.0.0", "NATS broker address.")
	flag.StringVar(&opts.NatsPort, "q", "4222", "NATS broker port.")
	flag.BoolVar(&opts.Version, "version", false, "Print version information.")
	flag.BoolVar(&opts.Version, "v", false, "Print version information.")
	flag.BoolVar(&opts.Help, "h", false, "Show help.")
	flag.BoolVar(&opts.Help, "help", false, "Show help.")

	flag.Parse()

	if opts.Version {
		tw := tabwriter.NewWriter(os.Stdout, 2, 1, 2, ' ', 0)
		fmt.Fprintf(tw, "Build Tag:    %s\n", Tag)
		fmt.Fprintf(tw, "Build Time:   %s\n", Time)
		fmt.Fprintf(tw, "Platform:     %s\n", Platform)
		fmt.Fprintf(tw, "Go Version:   %s\n", GoVersion)
		fmt.Fprintf(tw, "Build SHA-1:  %s\n", Revision)
		tw.Flush()
		os.Exit(0)
	}

	if opts.Help {
		fmt.Printf("%s\n", help)
		os.Exit(0)
	}

	// Connect to NATS broker
	if err := backoff.Retry(tryNatsConnect, backoff.NewExponentialBackOff()); err != nil {
		log.Fatalf("NATS: Can't connect: %v\n", err)
	} else {
		log.Println("OK")
	}

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
