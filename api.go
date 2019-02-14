package main

import "roccodev.pw/api/monthlies/profile"

func main() {
	/*
		prof := leaderboard.BedMonthlyLeaderboard{}
		prof.Get(0, 1)
		fmt.Println(prof.Leaderboard["e891f022633a4bf28f4030493fda9bbb"].Place)
	*/

	profile.GetLeaderboard()
}
