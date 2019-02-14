package winstreaks

import (
	"encoding/json"
	"roccodev.pw/api/url"
)

func GetLeaderboard(game string, from int, to int) map[string]WinstreakProfile {
	body := url.DownloadJson(url.StreaksLeaderboard(game, false, from, to))
	var res map[string]WinstreakProfile
	json.Unmarshal(body, &res)
	return res
}

func GetHistoricalLeaderboard(game string, from int, to int) map[string]WinstreakProfile {
	body := url.DownloadJson(url.StreaksLeaderboard(game, true, from, to))
	var res map[string]WinstreakProfile
	json.Unmarshal(body, &res)
	return res
}
