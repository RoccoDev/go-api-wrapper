package winstreaks

import (
	"encoding/json"
	"fmt"
	"roccodev.pw/api/url"
)

type WinstreakProfile struct {
	Username  string `json:"name"`
	Place     int
	Winstreak int
}

func (prof *WinstreakProfile) getFromRaw(data []byte) {
	e := json.Unmarshal(data, prof)
	if e != nil {
		fmt.Printf("Error occurred while parsing JSON: %s", e.Error())
	}
}

func (prof *WinstreakProfile) Get(game string, uuid string) {
	body := url.DownloadJson(url.StreaksProfile(game, false, uuid))
	prof.getFromRaw(body)
}

func (prof *WinstreakProfile) GetHistorical(game string, uuid string) {
	body := url.DownloadJson(url.StreaksProfile(game, true, uuid))
	prof.getFromRaw(body)
}
