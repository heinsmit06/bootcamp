package bootcamp

// import "fmt"

// func main() {
//     fmt.Println(RotN("salem", 1))
//     fmt.Println(RotN("abc", 13))
// }

func RotN(s string, n int) string {
	shifted := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char >= 'a' && char <= 'z' {
			shifted[i] = 'a' + (char-'a'+byte(n))%26
		} else if char >= 'A' && char <= 'Z' {
			shifted[i] = 'A' + (char-'A'+byte(n))%26
		} else {
			shifted[i] = char
		}
	}
	return string(shifted)
}
