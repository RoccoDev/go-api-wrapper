package profile

import (
	"encoding/json"
	"fmt"
	"roccodev.pw/api/url"
)

// Hide and Seek monthly profile.
type HideMonthlyProfile struct {
	MonthlyProfile
	HiderKills  int `json:"hider_kills"`
	SeekerKills int `json:"seeker_kills"`
}

func (prof *HideMonthlyProfile) getFromRaw(data []byte) {
	e := json.Unmarshal(data, prof)
	if e != nil {
		fmt.Printf("Error occurred while parsing JSON: %s", e.Error())
	}
}

func (prof *HideMonthlyProfile) Get(uuid string) {
	body := downloadJson(url.MonthlyProfile(url.UrlHide, uuid))
	prof.getFromRaw(body)
}
