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
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/cisco/senml"
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
	// Validate Content-Type
	ctype, err := resolveContentType(r.Header.Get("Content-Type"))
	if err != nil {
		// Return content type error
		w.WriteHeader(http.StatusUnsupportedMediaType)
		str := `{"response": "` + err.Error() + `"}`
		io.WriteString(w, str)
		return
	}
	//If is senML validate it
	if ctype == "senml+json" {
		var err error
		if _, err = senml.Decode(data, senml.JSON); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			str := `{"response": "senML validation faild"}`
			io.WriteString(w, str)
			return
		}
	}

	cid := bone.GetValue(r, "channel_id")

	// Publisher ID header
	hdr := r.Header.Get("Client-ID")

	// Publish message on MQTT via NATS
	m := NatsMsg{}
	m.Channel = cid
	m.Publisher = hdr
	m.Protocol = "http"
	m.Payload = data
	// Add ContentType to NatsMsg
	m.ContentType = ctype

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

func resolveContentType(content string) (result string, err error) {

	switch content {
	case "application/senml+json":
		return "senml+json", nil
	case "application/octet-stream":
		return "octet-stream", nil
	}
	return "", errors.New("Unsuproted Content-Type")

}
