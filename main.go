package main

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {

}

func main() {

}

/*
func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)
	payload := []byte(r.FormValue("payload"))

	event := Plex2U.PlexEvent{}
	json.Unmarshal(payload, &event)
	if !Plex2U.PlayerIsAllowed(event.Player.Uuid) {
		return
	}
	log.WithFields(map[string]interface{}{
		"event": event.Event,
		"uuid":  event.Player.Uuid,
		"title": event.Player.Title,
	}).Info("received " + event.Event + " event")
	if Plex2U.IsTimerEnabled() {
		if !Plex2U.IsItCurrentlyDayTime() {
			event.ForwardToIFTTT(payload)
			return
		}
		log.Debug("It's currently day time")
		return
	}
	event.ForwardToIFTTT(payload)

}

func main() {
	Plex2U.Configure()
	log.Info("starting server on port 8000")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
*/
