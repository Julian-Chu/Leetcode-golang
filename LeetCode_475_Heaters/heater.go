package leetcode475

import "math"

func findRadius(houses []int, heaters []int) int {

	// // dist := make([]int, len(heaters)+1)
	// lenOfHouse := len(houses)
	// lenOfHeaters := len(heaters)
	// radius := 0
	// if heaters[0] - houses[0]
	// 	radius = heaters[0] - houses[0]
	// }
	// if radius < (houses[lenOfHouse-1] - heaters[lenOfHeaters-1]) {
	// 	radius = houses[lenOfHouse-1] - heaters[lenOfHeaters-1]
	// }
	// // dist = append(dist, heaters[0]-house[0])
	// // dist = append(dist, house[lenOfHouse-1]-heaters[lenOfHeaters-1])

	// for i := 1; i < lenOfHeaters; i++ {
	// 	// dist = append(dist, (heaters[i]-heaters[i-1])/2)
	// 	if radius < (heaters[i]-heaters[i-1])/2 {
	// 		radius = (heaters[i] - heaters[i-1]) / 2
	// 	}
	// }
	radius := 0
	for i := 0; i < len(houses); i++ {
		house := houses[i]
		dist := math.Abs(float64(house - heaters[0]))
		for j := 1; j < len(heaters); j++ {
			if dist > math.Abs(float64(house-heaters[j])) {
				dist = math.Abs(float64(house - heaters[j]))
			}
		}
		if int(dist) > radius {
			radius = int(dist)
		}
	}
	return radius
}
