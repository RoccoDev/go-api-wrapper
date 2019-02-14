package url

import (
	"strconv"
	"strings"
)

// Main URL
const MainURL = "https://api.roccodev.pw/"

const MonthliesURL = "/monthlies/"
const ProfileURL = "/profile/"
const LeaderboardURL = "/leaderboard/"

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

func generateMonthlyURL(game string, endpoint string) strings.Builder {
	var builder strings.Builder
	builder.WriteString(MainURL)
	builder.WriteString(game)
	builder.WriteString(MonthliesURL)
	builder.WriteString(endpoint)
	return builder
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
