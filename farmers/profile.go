package farmers

import (
	"encoding/json"
	"fmt"
	"roccodev.pw/api/url"
)

type KillfarmerProfile struct {
	Username   string
	Place      int
	Violations int `json:"vl"`
	KPGRecord  int `json:"record"`
}

func (prof *KillfarmerProfile) getFromRaw(data []byte) {
	e := json.Unmarshal(data, prof)
	if e != nil {
		fmt.Printf("Error occurred while parsing JSON: %s", e.Error())
	}
}

func (prof *KillfarmerProfile) Get(uuid string) {
	body := url.DownloadJson(url.FarmersProfile(uuid))
	prof.getFromRaw(body)
}
