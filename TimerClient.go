package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

type SunsetSunrise struct {
	Results struct {
		Sunrise string `json:"sunrise"`
		Sunset  string `json:"sunset"`
	} `json:"results"`
	Status string `json:"status"`
}

func (s SunsetSunrise) GetSunrise() time.Time {
	time, _ := time.Parse(time.RFC3339, s.Results.Sunrise)
	return time
}
func (s SunsetSunrise) GetSunset() time.Time {
	time, _ := time.Parse(time.RFC3339, s.Results.Sunset)
	return time
}

func (s SunsetSunrise) InDayRange(current time.Time) bool {
	return current.Unix() > s.GetSunrise().Unix() && current.Unix() < s.GetSunset().Unix()
}

func IsItCurrentlyDayTime() bool {

	lat, lng := GetLocation()
	resp, err := http.Get("https://api.sunrise-sunset.org/json?lat=" + lat + "&lng=" + lng + "&formatted=0")
	if err != nil {
		return false
	}
	body, _ := ioutil.ReadAll(resp.Body)
	sunsetTime := SunsetSunrise{}
	json.Unmarshal(body, &sunsetTime)
	log.WithFields(logrus.Fields{
		"sunset":   sunsetTime.GetSunset(),
		"sunrise":  sunsetTime.GetSunrise(),
		"Daylight": sunsetTime.InDayRange(time.Now()),
	}).Debug("Checked if current day is on")
	return sunsetTime.InDayRange(time.Now())
}
