# Go API Wrapper for roccodev.pw

## Usage
### Monthlies
#### Profile
Instantiate a `MonthlyProfile` struct and populate it with the `Get` method.  
Example:
```go
import "roccodev.pw/api/monthlies/profile"

...

var prof profile.BedMonthlyProfile
prof.Get(/* RoccoDev's UUID */ "bba224a20bff4913b04227ca3b60973f")

fmt.Println(prof.Beds)
```

#### Leaderboard
You can use the `GetXXXLeaderboard` methods to get a `map[string]MonthlyProfile`.  
Example:
```go
import "roccodev.pw/api/monthlies/leaderboard"

...

lb := leaderboard.GetBedwarsLeaderboard(/* Start */ 0, /* End */ 1)
```

### BED Kill farmers
Use it the same way as the monthlies, the struct is called `KillfarmerProfile`.

### Coverage  
- [x] Monthlies
- [ ] Winstreaks
- [x] BED Kill farmers