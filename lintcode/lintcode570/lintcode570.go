package lintcode570

import "strconv"

func findMissing2(n int, str string) int {
	visited := make(map[int]bool)
	for i := 1; i < n+1; i++ {
		visited[i] = false
	}

	dfs(str, 0, visited, n)
	for key := range visited {
		if !visited[key] {
			return key
		}
	}
	return -1
}

func dfs(str string, start int, visited map[int]bool, n int) bool {
	if start >= len(str) {
		return true
	}

	if str[start] == '0' {
		return false
	}

	num, _ := strconv.Atoi(str[start : start+1])
	if num > 0 && num <= n && visited[num] == false {
		visited[num] = true
		if dfs(str, start+1, visited, n) {
			return true
		}
		visited[num] = false
	}

	if start+1 >= len(str) {
		return false
	}
	num, _ = strconv.Atoi(str[start : start+2])
	if num > 0 && num <= n && visited[num] == false {
		visited[num] = true
		if dfs(str, start+2, visited, n) {
			return true
		}
		visited[num] = false
	}
	return false
}
