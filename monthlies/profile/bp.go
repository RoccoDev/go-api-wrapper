package profile

import (
	"encoding/json"
	"fmt"
	"roccodev.pw/api/url"
)

// BlockParty monthly profile.
type BpMonthlyProfile struct {
	MonthlyProfile
	Placings     int
	Eliminations int
}

func (prof *BpMonthlyProfile) getFromRaw(data []byte) {
	e := json.Unmarshal(data, prof)
	if e != nil {
		fmt.Printf("Error occurred while parsing JSON: %s", e.Error())
	}
}

func (prof *BpMonthlyProfile) Get(uuid string) {
	body := downloadJson(url.MonthlyProfile(url.UrlBp, uuid))
	prof.getFromRaw(body)
}
