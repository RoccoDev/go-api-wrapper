package leaderboard

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"roccodev.pw/api/monthlies/profile"
)

func GetLeaderboard() {

	body := downloadJson("https://api.roccodev.pw/bed/monthlies/leaderboard?from=0&to=1")

	var res map[string]profile.BedMonthlyProfile
	json.Unmarshal(body, &res)

	println(res["e891f022633a4bf28f4030493fda9bbb"].Place)

}

func downloadJson(url string) []byte {
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
