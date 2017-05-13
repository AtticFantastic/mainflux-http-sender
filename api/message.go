/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-zoo/bone"
)

// sendMessage function
func sendMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	if len(data) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		str := `{"response": "no data provided"}`
		io.WriteString(w, str)
		return
	}

	//TODO Handle Content-Type
	//TODO If is senML validate it
	//TODO Add ContentType to NatsMsg

	cid := bone.GetValue(r, "channel_id")

	// Publisher ID header
	hdr := r.Header.Get("Client-ID")

	// Publish message on MQTT via NATS
	m := NatsMsg{}
	m.Channel = cid
	m.Publisher = hdr
	m.Protocol = "http"
	m.Payload = data

	b, err := json.Marshal(m)
	if err != nil {
		log.Print(err)
	}
	NatsConn.Publish("msg.http", b)

	// Send back response to HTTP client
	// We have accepted the request and published it over MQTT,
	// but we do not know if it will be executed or not (MQTT is not req-reply protocol)
	w.WriteHeader(http.StatusAccepted)
	str := `{"response": "message sent"}`
	io.WriteString(w, str)
}
