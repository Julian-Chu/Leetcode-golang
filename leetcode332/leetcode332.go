package leetcode332

func findItinerary(tickets [][]string) []string {
	inout := make(map[string]int)
	for _, ticket := range tickets {
		inout[ticket[0]]++
		inout[ticket[1]]--
	}

	dest := "JFK"
	for key, val := range inout {
		if val == -1 {
			dest = key
		}
	}

	flags := make([]bool, len(tickets))
	res := make([]string, 0)
	var dfs func(curTicket []string, restSteps int, path []string)
	dfs = func(curTicket []string, restSteps int, path []string) {
		to := curTicket[1]
		path = path[:len(path):len(path)]
		path = append(path, to)
		if restSteps == 0 {
			if to != dest {
				return
			}
			if len(res) == 0 {
				res = path
			} else {
			loop:
				for i := 0; i < len(res); i++ {
					for j := 0; j < 3; j++ {
						if path[i][j] == res[i][j] {
							continue
						}
						if path[i][j] < res[i][j] {
							res = path
						}
						break loop
					}
				}
			}
			return
		}

		for i, ticket := range tickets {
			if flags[i] == true {
				continue
			}
			from := ticket[0]
			if to != from {
				continue
			}
			restSteps--
			flags[i] = true
			dfs(ticket, restSteps, path)
			flags[i] = false
			restSteps++
		}
	}

	for _, ticket := range tickets {
		if ticket[0] != "JFK" {
			continue
		}

		path := []string{"JFK"}
		restSteps := len(tickets) - 1
		dfs(ticket, restSteps, path)
	}
	return res
}
