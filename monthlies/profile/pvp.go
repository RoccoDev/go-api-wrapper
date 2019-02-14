package profile

import (
	"encoding/json"
	"fmt"
	"roccodev.pw/api/url"
)

type pvpMonthlyProfile struct {
	MonthlyProfile
	Kills  int
	Deaths int
}

type GntMonthlyProfile struct {
	pvpMonthlyProfile
}

type GntmMonthlyProfile struct {
	GntMonthlyProfile
}

type DrMonthlyProfile struct {
	pvpMonthlyProfile
}

type SkyMonthlyProfile struct {
	pvpMonthlyProfile
}

func (prof *pvpMonthlyProfile) getFromRaw(data []byte) {
	e := json.Unmarshal(data, prof)
	if e != nil {
		fmt.Printf("Error occurred while parsing JSON: %s", e.Error())
	}
}

func (prof *pvpMonthlyProfile) Get(uuid string) {
	body := downloadJson(url.MonthlyProfile(url.UrlSkywars, uuid))
	prof.getFromRaw(body)
}

func (prof *SkyMonthlyProfile) Get(uuid string) {
	body := downloadJson(url.MonthlyProfile(url.UrlSkywars, uuid))
	prof.getFromRaw(body)
}

func (prof *GntMonthlyProfile) Get(uuid string) {
	body := downloadJson(url.MonthlyProfile(url.UrlGnt, uuid))
	prof.getFromRaw(body)
}

func (prof *GntmMonthlyProfile) Get(uuid string) {
	body := downloadJson(url.MonthlyProfile(url.UrlGntm, uuid))
	prof.getFromRaw(body)
}

func (prof *DrMonthlyProfile) Get(uuid string) {
	body := downloadJson(url.MonthlyProfile(url.UrlDr, uuid))
	prof.getFromRaw(body)
}
