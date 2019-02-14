package url

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// Main URL
const MainURL = "https://api.roccodev.pw/"

const MonthliesURL = "/monthlies"
const FarmersURL = "/farmers"
const WinstreaksURL = "/winstreaks"
const HistoURL = "/winstreaks/historical"

const ProfileURL = "/profile/"
const LeaderboardURL = "/leaderboard/"

const BedSolo = "beds"
const BedTeams = "bedt"
const BedLTM = "bedx"
const BedDuos = "bedd"

/* Game names */
const UrlBedwars = "bed"
const UrlSkywars = "sky"
const UrlHide = "hide"
const UrlDr = "dr"
const UrlBp = "bp"
const UrlCai = "cai"
const UrlGnt = "gnt"
const UrlGntm = "gntm"
const UrlTimv = "timv"

func generateMonthlyURL(game string, endpoint string) *strings.Builder {
	var builder strings.Builder
	builder.WriteString(MainURL)
	builder.WriteString(game)
	builder.WriteString(MonthliesURL)
	builder.WriteString(endpoint)
	return &builder
}

func generateStreaksUrl(game string, endpoint string, historical bool) *strings.Builder {
	var builder strings.Builder
	builder.WriteString(MainURL)
	builder.WriteString(game)
	if historical {
		builder.WriteString(HistoURL)
	} else {
		builder.WriteString(WinstreaksURL)
	}
	builder.WriteString(endpoint)
	return &builder
}

func generateFarmersUrl(endpoint string) *strings.Builder {
	var builder strings.Builder
	builder.WriteString(MainURL)
	builder.WriteString(UrlBedwars)
	builder.WriteString(FarmersURL)
	builder.WriteString(endpoint)
	return &builder
}

func MonthlyProfile(game string, uuid string) string {
	builder := generateMonthlyURL(game, ProfileURL)
	builder.WriteString(uuid)
	return builder.String()
}

func MonthlyLeaderboard(game string, from int, to int) string {
	builder := generateMonthlyURL(game, LeaderboardURL)
	builder.WriteString("?from=")
	builder.WriteString(strconv.Itoa(from))
	builder.WriteString("&to=")
	builder.WriteString(strconv.Itoa(to))
	return builder.String()
}

func FarmersProfile(uuid string) string {
	builder := generateFarmersUrl(ProfileURL)
	builder.WriteString(uuid)
	return builder.String()
}

func FarmersLeaderboard(from int, to int, order string) string {
	builder := generateFarmersUrl(LeaderboardURL)
	builder.WriteString("?from=")
	builder.WriteString(strconv.Itoa(from))
	builder.WriteString("&to=")
	builder.WriteString(strconv.Itoa(to))
	builder.WriteString("&order=")
	builder.WriteString(order)
	return builder.String()
}

func StreaksProfile(game string, historical bool, uuid string) string {
	builder := generateStreaksUrl(game, ProfileURL, historical)
	builder.WriteString(uuid)
	return builder.String()
}

func StreaksLeaderboard(game string, historical bool, from int, to int) string {
	builder := generateStreaksUrl(game, LeaderboardURL, historical)
	builder.WriteString("?from=")
	builder.WriteString(strconv.Itoa(from))
	builder.WriteString("&to=")
	builder.WriteString(strconv.Itoa(to))
	return builder.String()
}

func DownloadJson(url string) []byte {
	res, e := http.Get(url)
	if e != nil {
		fmt.Printf("Error occurred while downloading data: %s", e.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("Error occurred while parsing body: %s", err.Error())
	}

	return []byte(string(body))

}
