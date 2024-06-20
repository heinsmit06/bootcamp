package bootcamp

func ReverseBits(n byte) byte {
	var reversed byte
	for i := 0; i < 8; i++ {
		bit := (n >> i) & 1        // Extract the i-th bit from n
		reversed |= bit << (7 - i) // Set the corresponding bit in reversed
	}
	return reversed
}
