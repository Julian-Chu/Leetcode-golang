package LeetCode_841_KeysAndRooms

func canVisitAllRooms_bfs(rooms [][]int) bool {
	if len(rooms) == 0 {
		return false
	}

	queue := []int{0}

	visited := make([]bool, len(rooms))
	for len(queue) > 0 {
		size := len(queue)
		for j := 0; j < size; j++ {
			roomNum := queue[j]
			if visited[roomNum] {
				continue
			}
			visited[roomNum] = true
			queue = append(queue, rooms[roomNum]...)
		}

		allVisited := true
		for i := range visited {
			if visited[i] {
				continue
			}
			allVisited = false
			break
		}

		if allVisited {
			return true
		}

		queue = queue[size:]
	}
	return false
}

func canVisitAllRooms_dfs(rooms [][]int) bool {
	n := len(rooms)
	if n == 0 {
		return false
	}

	var dfs func(next int, vistited []bool)
	dfs = func(key int, visited []bool) {
		if visited[key] {
			return
		}
		visited[key] = true
		keys := rooms[key]
		for _, key := range keys {
			dfs(key, visited)
		}
	}
	visited := make([]bool, n)
	dfs(0, visited)
	for i := range visited {
		if visited[i] == false {
			return false
		}
	}
	return true
}
