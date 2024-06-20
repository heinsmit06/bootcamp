package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func main() {
	args := os.Args
	seconds := Atoi1(args[1])

	if seconds > 2147483647 || seconds < 0 {
		ap.PutRune('N')
		ap.PutRune('V')
	}

	days, hours, minutes := 0, 0, 0

	if seconds > 86400 {
		days = days + seconds/86400
		seconds = seconds % 86400
	}

	if seconds > 3600 {
		hours += seconds / 3600
		seconds = seconds % 3600
	}

	if seconds > 60 {
		minutes += seconds / 60
		seconds = seconds % 60
	}

	if days != 0 {
		PrintString(Itoa1(days) + "d")
		ap.PutRune(' ')
	}
	if hours != 0 {
		PrintString(Itoa1(hours) + "h")
		ap.PutRune(' ')
	}
	if minutes != 0 {
		PrintString(Itoa1(minutes) + "m")
		ap.PutRune(' ')
	}
	if seconds != 0 {
		PrintString(Itoa1(seconds) + "s")
	}

	ap.PutRune('\n')
}

func PrintString(s string) {
	for i := 0; i < len(s); i++ {
		ap.PutRune(rune(s[i]))
	}
}

func Itoa1(n int) string {
	if n == 0 {
		return "0"
	}

	var result string
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}

	for n > 0 {
		digit := n % 10
		result = string(rune(digit+48)) + result
		n /= 10
	}

	return sign + result
}

func Atoi1(s string) int {
	var res int
	power := 0
	for i := len(s) - 1; i >= 0; i-- {
		res += int(s[i]-48) * PowRecursion1(10, power)
		power++
	}

	return res
}

func PowRecursion1(x int, power int) int {
	if power < 0 {
		return -1
	} else if power == 0 {
		return 1
	}

	var result int = 1
	if power > 0 {
		result = x * PowRecursion1(x, power-1)
	}
	return result
}
