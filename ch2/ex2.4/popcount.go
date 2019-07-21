package popcount

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	count := 0
	var mask uint64 = 1
	for i := 0; i < 64; i++ {
		if x&mask > 0 {
			count++
		}
		x >>= 1
	}
	return count
}
