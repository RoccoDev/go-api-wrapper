package profile

import (
	"encoding/json"
	"fmt"
	"roccodev.pw/api/url"
)

type TimvMonthlyProfile struct {
	MonthlyProfile
	MostChange      string `json:"most_change"`
	DetectivePoints int    `json:"d_points"`
	TraitorPoints   int    `json:"t_points"`
	InnocentPoints  int    `json:"i_points"`
}

func (prof *TimvMonthlyProfile) getFromRaw(data []byte) {
	e := json.Unmarshal(data, prof)
	if e != nil {
		fmt.Printf("Error occurred while parsing JSON: %s", e.Error())
	}
}

func (prof *TimvMonthlyProfile) Get(uuid string) {
	body := downloadJson(url.MonthlyProfile(url.UrlTimv, uuid))
	prof.getFromRaw(body)
}
