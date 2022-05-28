package LeetCode_1152_AnalyzeUserWebsiteVisitPattern

import (
	"sort"
	"strings"
)

func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
	type visit struct {
		website   string
		timestamp int
	}
	users := make(map[string][]visit, 50)

	for i := 0; i < len(username); i++ {
		name := username[i]
		ts := timestamp[i]
		w := website[i]

		if _, ok := users[name]; !ok {
			users[name] = make([]visit, 0, 50)
		}
		users[name] = append(users[name], visit{
			website:   w,
			timestamp: ts,
		})
	}

	scores := make(map[[3]string]int, 50)
	for _, visits := range users {
		sort.Slice(visits, func(i, j int) bool {
			return visits[i].timestamp < visits[j].timestamp
		})

		userUniquePatterns := map[[3]string]struct{}{}
		for i := 0; i < len(visits)-2; i++ {
			for j := i + 1; j < len(visits)-1; j++ {
				for k := j + 1; k < len(visits); k++ {
					pattern := [3]string{
						visits[i].website,
						visits[j].website,
						visits[k].website,
					}
					if _, ok := userUniquePatterns[pattern]; !ok {
						userUniquePatterns[pattern] = struct{}{}
					}
				}
			}
		}

		for p := range userUniquePatterns {
			if _, ok := scores[p]; !ok {
				scores[p] = 0
			}
			scores[p]++
		}
	}

	maxScore := 0
	var maxPattern [3]string
	for pattern, score := range scores {
		if score > maxScore {
			maxScore = score
			maxPattern = pattern
			continue
		}
		if score == maxScore {
			for i := 0; i < 3; i++ {
				if strings.Compare(pattern[i], maxPattern[i]) == 0 {
					continue
				}

				if strings.Compare(pattern[i], maxPattern[i]) == -1 {
					maxPattern = pattern
				}
				break
			}
		}
	}
	return maxPattern[:]
}
