package lintcode128

/**
 * @param key: A string you should hash
 * @param HASH_SIZE: An integer
 * @return: An integer
 */
func hashCode(key string, HASH_SIZE int) int {
	ans := 0
	for i := range key {
		ans = (ans*33 + int(key[i])) % HASH_SIZE
	}
	return ans

}
