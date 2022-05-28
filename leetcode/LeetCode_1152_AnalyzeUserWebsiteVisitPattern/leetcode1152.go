package LeetCode_1152_AnalyzeUserWebsiteVisitPattern

import "sort"

type visit struct {
	timestamp int
	website   string
}

func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
	// website+timestamp by username
	// sort

	// usernames hashsets by each triset
	// populate by visiting all users

	visits := make(map[string][]visit, 0)
	for i := 0; i < len(username); i++ {
		visits[username[i]] = append(visits[username[i]], visit{timestamp[i], website[i]})
	}

	for _, userVisits := range visits {
		sort.Slice(userVisits, func(i, j int) bool {
			return userVisits[i].timestamp < userVisits[j].timestamp
		})
	}

	patterns := make(map[[3]string]map[string]struct{})
	for user, userVisits := range visits {
		if len(userVisits) < 3 {
			continue
		}
		// todo: have to generate all subsequences
		for i := 0; i < len(userVisits)-2; i++ {
			for j := i + 1; j < len(userVisits)-1; j++ {
				for k := j + 1; k < len(userVisits); k++ {
					pattern := [3]string{userVisits[i].website, userVisits[j].website, userVisits[k].website}
					if patterns[pattern] == nil {
						patterns[pattern] = make(map[string]struct{})
					}
					patterns[pattern][user] = struct{}{}
				}
			}
		}
	}

	var res [3]string
	var maxUsers int

	for pattern, users := range patterns {
		if len(users) > maxUsers {
			maxUsers = len(users)
			res = pattern
		} else if len(users) == maxUsers && less(pattern, res) {
			res = pattern
		}
	}

	return []string{res[0], res[1], res[2]}
}

func less(p1, p2 [3]string) bool {
	if p1[0] == p2[0] {
		if p1[1] == p2[1] {
			return p1[2] < p2[2]
		}
		return p1[1] < p2[1]
	}
	return p1[0] < p2[0]
}

//func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
//	type visit struct {
//		website   string
//		timestamp int
//	}
//	users := make(map[string][]visit, 50)
//
//	for i := 0; i < len(username); i++ {
//		name := username[i]
//		ts := timestamp[i]
//		w := website[i]
//
//		if _, ok := users[name]; !ok {
//			users[name] = make([]visit, 0, 50)
//		}
//		users[name] = append(users[name], visit{
//			website:   w,
//			timestamp: ts,
//		})
//	}
//
//	scores := make(map[[3]string]int, 50)
//	for _, visits := range users {
//		sort.Slice(visits, func(i, j int) bool {
//			return visits[i].timestamp < visits[j].timestamp
//		})
//
//		userUniquePatterns := map[[3]string]struct{}{}
//		for i := 0; i < len(visits)-2; i++ {
//			for j := i + 1; j < len(visits)-1; j++ {
//				for k := j + 1; k < len(visits); k++ {
//					pattern := [3]string{
//						visits[i].website,
//						visits[j].website,
//						visits[k].website,
//					}
//					if _, ok := userUniquePatterns[pattern]; !ok {
//						userUniquePatterns[pattern] = struct{}{}
//					}
//				}
//			}
//		}
//
//		for p := range userUniquePatterns {
//			if _, ok := scores[p]; !ok {
//				scores[p] = 0
//			}
//			scores[p]++
//		}
//	}
//
//	maxScore := 0
//	var maxPattern [3]string
//	for pattern, score := range scores {
//		if score > maxScore {
//			maxScore = score
//			maxPattern = pattern
//			continue
//		}
//		if score == maxScore {
//			for i := 0; i < 3; i++ {
//				if strings.Compare(pattern[i], maxPattern[i]) == 0 {
//					continue
//				}
//
//				if strings.Compare(pattern[i], maxPattern[i]) == -1 {
//					maxPattern = pattern
//				}
//				break
//			}
//		}
//	}
//	return maxPattern[:]
//}
