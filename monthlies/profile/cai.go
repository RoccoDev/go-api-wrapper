package profile

import (
	"encoding/json"
	"fmt"
	"roccodev.pw/api/url"
)

type CaiMonthlyProfile struct {
	MonthlyProfile
	Captures int
	Caught   int
}

func (prof *CaiMonthlyProfile) getFromRaw(data []byte) {
	e := json.Unmarshal(data, prof)
	if e != nil {
		fmt.Printf("Error occurred while parsing JSON: %s", e.Error())
	}
}

func (prof *CaiMonthlyProfile) Get(uuid string) {
	body := downloadJson(url.MonthlyProfile(url.UrlCai, uuid))
	prof.getFromRaw(body)
}
