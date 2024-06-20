//package main
//
//import (
//	"os"
//
//	"github.com/alem-platform/ap"
//)
//
//func main() {
//	args := os.Args
//	seconds := Atoi1(args[1])
//
//	if seconds > 2147483647 || seconds < 0 {
//		ap.PutRune('N')
//		ap.PutRune('V')
//
//	}
//
//	days, hours, minutes := 0, 0, 0
//
//	if seconds > 86400 {
//		days = days + seconds/86400
//		seconds = seconds % 86400
//	}
//
//	if seconds > 3600 {
//		hours += seconds / 3600
//		seconds = seconds % 3600
//	}
//
//	if seconds > 60 {
//		minutes += seconds / 60
//		seconds = seconds % 60
//	}
//
//	if days != 0 {
//		PrintString(Itoa1(days) + "d")
//		ap.PutRune(' ')
//	}
//	if hours != 0 {
//		PrintString(Itoa1(hours) + "h")
//		ap.PutRune(' ')
//	}
//	if minutes != 0 {
//		PrintString(Itoa1(minutes) + "m")
//		ap.PutRune(' ')
//	}
//	if seconds != 0 {
//		PrintString(Itoa1(seconds) + "s")
//	}
//
//	ap.PutRune('\n')
//}
//
//func PrintString(s string) {
//	for i := 0; i < len(s); i++ {
//		ap.PutRune(rune(s[i]))
//	}
//}
//
//func Itoa1(n int) string {
//	if n == 0 {
//		return "0"
//	}
//
//	var result string
//	sign := ""
//	if n < 0 {
//		sign = "-"
//		n = -n
//	}
//
//	for n > 0 {
//		digit := n % 10
//		result = string(rune(digit+48)) + result
//		n /= 10
//	}
//
//	return sign + result
//}
//
//func Atoi1(s string) int {
//	var res int
//	power := 0
//	for i := len(s) - 1; i >= 0; i-- {
//		res += int(s[i]-48) * PowRecursion1(10, power)
//		power++
//	}
//
//	return res
//}
//
//func PowRecursion1(x int, power int) int {
//	if power < 0 {
//		return -1
//	} else if power == 0 {
//		return 1
//	}
//
//	var result int = 1
//	if power > 0 {
//		result = x * PowRecursion1(x, power-1)
//	}
//	return result
//}

package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func Atoi(s string) int {
	var num int64
	var sign int64 = 1
	var start int = 0

	if len(s) == 0 {
		return 0
	}

	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		num = num*10 + int64(s[i]-'0')
	}

	result := sign * num

	if result < -2147483648 || result > 2147483647 {
		return 0
	}

	return int(result)
}

func main() {
	if len(os.Args) != 2 {
		ap.PutRune('N')
		ap.PutRune('V')
		ap.PutRune('\n')
		return
	}

	input := Atoi(os.Args[1])

	if input == 0 {
		ap.PutRune('N')
		ap.PutRune('V')
		ap.PutRune('\n')
		return
	}

	if input < 0 {
		ap.PutRune('N')
		ap.PutRune('V')
		ap.PutRune('\n')
		return
	}

	days := input / 86400
	input %= 86400
	hours := input / 3600
	input %= 3600
	minutes := input / 60
	seconds := input % 60

	if days > 0 {
		recursiveConv(days)
		ap.PutRune('d')
		ap.PutRune(' ')
	}
	if hours > 0 {
		recursiveConv(hours)
		ap.PutRune('h')
		ap.PutRune(' ')
	}
	if minutes > 0 {
		recursiveConv(minutes)
		ap.PutRune('m')
		ap.PutRune(' ')
	}
	if seconds > 0 || (days == 0 && hours == 0 && minutes == 0) {
		recursiveConv(seconds)
		ap.PutRune('s')
	}
	ap.PutRune('\n')
}

func recursiveConv(n int) {
	if n < 0 {
		ap.PutRune('-')
		n = -n
	}
	if n < 10 {
		ap.PutRune(rune(n + '0'))
		return
	}
	recursiveConv(n / 10)
	ap.PutRune(rune(n%10 + '0'))
}
