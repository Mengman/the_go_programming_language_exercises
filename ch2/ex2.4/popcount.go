package popcount

// Exercise 2.4: Write a version of PopCount that counts bits by shifting its argument through 64
// bit positions, testing the rightmost bit each time. Compare its performance to the table-lookup version.

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
