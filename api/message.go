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
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/cisco/senml"
	"github.com/go-zoo/bone"
)

const (
	senMl string = "application/senml+json"
	blob  string = "application/octet-stream"
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
	ctype, err := validateContentType(r.Header.Get("Content-Type"))
	if err != nil {
		// Return content type error
		w.WriteHeader(http.StatusUnsupportedMediaType)
		str := `{"response": "` + err.Error() + `"}`
		io.WriteString(w, str)
		return
	}
	//If is senML validate it
	if ctype == senMl {
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
	m.ContentType = cleanContentType(ctype)

	fmt.Printf(m.ContentType)

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

func cleanContentType(ct string) string {
	return strings.Split(ct, "/")[1]
}

func validateContentType(content string) (result string, err error) {

	switch content {
	case blob, senMl:
		return content, nil
	}
	return content, errors.New("Unsuproted Content-Type")

}
