package leaderboard

import (
	"encoding/json"
	"roccodev.pw/api/monthlies/profile"
	"roccodev.pw/api/url"
)

func GetBedwarsLeaderboard(from int, to int) map[string]profile.BedMonthlyProfile {
	body := downloadJson(url.MonthlyLeaderboard(url.UrlBedwars, from, to))
	var res map[string]profile.BedMonthlyProfile
	json.Unmarshal(body, &res)
	return res
}

func GetSkywarsLeaderboard(from int, to int) map[string]profile.SkyMonthlyProfile {
	body := downloadJson(url.MonthlyLeaderboard(url.UrlSkywars, from, to))
	var res map[string]profile.SkyMonthlyProfile
	json.Unmarshal(body, &res)
	return res
}

func GetCaiLeaderboard(from int, to int) map[string]profile.CaiMonthlyProfile {
	body := downloadJson(url.MonthlyLeaderboard(url.UrlCai, from, to))
	var res map[string]profile.CaiMonthlyProfile
	json.Unmarshal(body, &res)
	return res
}

func GetDrLeaderboard(from int, to int) map[string]profile.DrMonthlyProfile {
	body := downloadJson(url.MonthlyLeaderboard(url.UrlDr, from, to))
	var res map[string]profile.DrMonthlyProfile
	json.Unmarshal(body, &res)
	return res
}

func GetGntLeaderboard(from int, to int) map[string]profile.GntMonthlyProfile {
	body := downloadJson(url.MonthlyLeaderboard(url.UrlGnt, from, to))
	var res map[string]profile.GntMonthlyProfile
	json.Unmarshal(body, &res)
	return res
}

func GetGntmLeaderboard(from int, to int) map[string]profile.GntmMonthlyProfile {
	body := downloadJson(url.MonthlyLeaderboard(url.UrlGntm, from, to))
	var res map[string]profile.GntmMonthlyProfile
	json.Unmarshal(body, &res)
	return res
}

func GetTimvLeaderboard(from int, to int) map[string]profile.TimvMonthlyProfile {
	body := downloadJson(url.MonthlyLeaderboard(url.UrlTimv, from, to))
	var res map[string]profile.TimvMonthlyProfile
	json.Unmarshal(body, &res)
	return res
}

func GetHideLeaderboard(from int, to int) map[string]profile.HideMonthlyProfile {
	body := downloadJson(url.MonthlyLeaderboard(url.UrlHide, from, to))
	var res map[string]profile.HideMonthlyProfile
	json.Unmarshal(body, &res)
	return res
}

func GetBpLeaderboard(from int, to int) map[string]profile.BpMonthlyProfile {
	body := downloadJson(url.MonthlyLeaderboard(url.UrlBp, from, to))
	var res map[string]profile.BpMonthlyProfile
	json.Unmarshal(body, &res)
	return res
}
