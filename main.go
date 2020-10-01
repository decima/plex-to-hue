package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

var log = logrus.New()

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)
	payload := []byte(r.FormValue("payload"))

	log.Debug("Json Received:", string(payload))
	event := PlexEvent{}
	json.Unmarshal(payload, &event)
	if !PlayerIsAllowed(event.Player.Uuid) {
		return
	}
	log.WithFields(map[string]interface{}{
		"event": event.Event,
		"uuid":  event.Player.Uuid,
		"title": event.Player.Title,
	}).Info("received " + event.Event + " event")
	if isTimerEnabled() {
		if !IsItCurrentlyDayTime() {
			event.ForwardToIFTTT(payload)
			return
		}
		log.Debug("It's currently day time")
		return
	}
	event.ForwardToIFTTT(payload)

}

func main() {
	Configure()
	log.Info("starting server on port 8000")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
