package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	count := byte(0)
	for i := uint8(0); i < 8; i++ {
		count += pc[byte(x>>(i*8))]
	}
	return int(count)
}
