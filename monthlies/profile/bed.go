package profile

import (
	"encoding/json"
	"fmt"
	"roccodev.pw/api/url"
)

// Bedwars monthly profile.
type BedMonthlyProfile struct {
	pvpMonthlyProfile
	Username string `json:"name"`
	Beds     int
	Teams    int
}

func (prof *BedMonthlyProfile) getFromRaw(data []byte) {
	e := json.Unmarshal(data, prof)
	if e != nil {
		fmt.Printf("Error occurred while parsing JSON: %s", e.Error())
	}
}

func (prof *BedMonthlyProfile) Get(uuid string) {
	body := downloadJson(url.MonthlyProfile(url.UrlBedwars, uuid))
	prof.getFromRaw(body)
}
