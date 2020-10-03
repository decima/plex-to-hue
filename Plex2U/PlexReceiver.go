package Plex2U

import (
	"bytes"
	"net/http"
)

type PlexEvent struct {
	Event string `json:"event"`
	Player struct{
		Title string `json:"title"`
		Uuid string `json:"uuid"`
	} `json:"player"`
}

func (PlexEvent PlexEvent) IsValid() bool {
	var AvailableEvents = [...]string{
		"library.on.deck",
		"library.new",
		"media.pause",
		"media.play",
		"media.rate",
		"media.resume",
		"media.scrobble",
		"media.stop",
		"admin.database.backup",
		"admin.database.corrupted",
		"device.new",
		"playback.started",
	}

	for _, b := range AvailableEvents {
		if b == PlexEvent.Event {
			return true
		}
	}
	return false
}
func (e PlexEvent) ForwardToIFTTT(body []byte) bool {
	if e.IsValid() {
		http.Post("https://maker.ifttt.com/trigger/"+e.Event+"/with/key/"+GetWebHook(), "Application/json", bytes.NewReader(body))
		return true
	}
	return false
}
