package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func main() {
	cw, cl, cb := false, false, false
	arg := os.Args[1]
	prm := ""
	filename := []string{}
	for i := 0; i < len(arg); i++ {
		if arg[0] == '-' {
			prm = os.Args[1]
			filename = os.Args[2:]
		} else {
			filename = os.Args[1:]
		}
	}
	for i := 0; i < len(filename); i++ {
		b, err := os.ReadFile(filename[i])
		if err != nil {
			Print("my_wc: ")
			Print(filename[i])
			Print(": No such file or directory\n")
			return
		}
		words := string(b)
		for _, value := range prm {
			if value == 'l' {
				cl = true
			}
			if value == 'w' {
				cw = true
			}
			if value == 'c' {
				cb = true
			}

		}
		if !cl && !cw && !cb {
			Print(Itoa(countnewline(words)))
			Print(" ")
			Print(Itoa(countwords(words)))
			Print(" ")
			Print(Itoa(countbytes(words)))
			Print(" ")
		}
		if cl {
			Print(Itoa(countnewline(words)))
			Print(" ")
		}
		if cw {
			Print(Itoa(countwords(words)))
			Print(" ")
		}
		if cb {
			Print(Itoa(countbytes(words)))
			Print(" ")
		}
		Print(filename[i])
		Print("\n")
	}
}

func countnewline(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 10 {
			count++
		}
	}

	return count
}

func countwords(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			count++
			for k := 0; k < len(s); k++ {
				if k+i == len(s)-1 {
					i = i + k
					break
				}
				if s[i+k] == '\n' {
					i = i + k
					break
				}
				if s[i+k] == 32 {
					i = i + k
					break
				}
			}
		}
	}
	return count
}

func countbytes(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 0 && s[i] <= 255 {
			count++
		}
	}
	return count
}

func Print(s string) {
	for _, value := range s {
		ap.PutRune(value)
	}
}

func Itoa(n int) string {
	var revconverted string
	var converted string
	digit := 0
	if n == 0 {
		return "0"
	}
	if n < 0 {
		converted = converted + "-"
		n *= -1
	}

	for n != 0 {
		digit = n % 10
		revconverted = revconverted + string(digit+'0')
		n /= 10
	}
	for j := len(revconverted) - 1; j >= 0; j-- {
		converted = converted + string(revconverted[j])
	}

	return converted
}
