/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package api

import (
	"log"

	"github.com/nats-io/go-nats"
)

type (
	NatsMsg struct {
		Channel     string `json:"channel"`
		Publisher   string `json:"publisher"`
		Protocol    string `json:"protocol"`
		Payload     []byte `json:"payload"`
		ContentType string `json:"content_type"`
	}
)

var (
	NatsConn *nats.Conn
)

func NatsInit(host string, port string) error {
	/** Connect to NATS broker */
	var err error
	NatsConn, err = nats.Connect("nats://" + host + ":" + port)
	if err != nil {
		log.Println(err)
	}

	return err
}
