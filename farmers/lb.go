package farmers

import (
	"encoding/json"
	"roccodev.pw/api/url"
)

const (
	OrderByKPGRecord  = "record"
	OrderByViolations = "vl"
)

func GetLeaderboard(from int, to int, order string) map[string]KillfarmerProfile {
	body := url.DownloadJson(url.FarmersLeaderboard(from, to, order))
	var res map[string]KillfarmerProfile
	json.Unmarshal(body, &res)
	return res
}
