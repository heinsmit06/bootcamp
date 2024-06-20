package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func printLine(s string) {
	for i := 0; i < len(s); i++ {
		ap.PutRune(rune(s[i]))
	}
}

func LenMaxLenth(content string) int {
	max := 0
	temp := 0
	for i := 0; i < len(content); i++ {
		temp++
		if (content[i] == '\n' || i == len(content)-1) && max < temp {
			if i == len(content)-1 && content[i] != '\n' {
				temp += 1
			}
			max = temp - 1
			temp = 0
		}
	}
	return max
}

func Top(size int) {
	ap.PutRune(' ')
	for i := 0; i < (size + 2); i++ {
		ap.PutRune('_')
	}
	ap.PutRune('\n')
}

func Bottom(size int) {
	ap.PutRune(' ')
	for i := 0; i < (size + 2); i++ {
		ap.PutRune('-')
	}
	ap.PutRune('\n')
}

func Strings(s string) []string {
	ans := []string{}
	j := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			ans = append(ans, s[j:i])
			j = i + 1
		}
	}
	ans = append(ans, s[j:len(s)])
	return ans
}

func TextBlock(content string) {
	maxLenth := LenMaxLenth(content)
	if maxLenth < 4 {
		maxLenth = 4
	}
	Top(maxLenth)
	lines := Strings(content)
	if len(lines) == 1 {
		printLine("< " + lines[0] + " >\n")
		Bottom(maxLenth)
		return
	}
	for i := 0; i < len(lines); i++ {
		var temp string
		if i == 0 {
			temp += "/ "
		} else if i == len(lines)-1 {
			temp += "\\ "
		} else {
			temp += "| "
		}
		temp += lines[i]
		for j := 0; j < maxLenth-len(lines[i]); j++ {
			temp += " "
		}
		if i == 0 {
			temp += " \\\n"
		} else if i == len(lines)-1 {
			temp += " /\n"
		} else {
			temp += " |\n"
		}
		printLine(temp)
	}
	Bottom(maxLenth)
}

func PrintCur(content string) {
	TextBlock(content)
	printLine("   \\\n")
	printLine("    \\ D\\___/\\\n")
	printLine("     \\ (0_o)\n")
	printLine("        (V)\n")
	printLine("----oOo--U--oOo------\n")
	printLine("_______|_______|_____\n")
}

func PrintMaltipoo(content string) {
	TextBlock(content)
	printLine("  \\\n")
	printLine("   \\ __    __\n")
	printLine("   o-''))_____\\\\\n")
	printLine("   \"--__/ * * * )\n")
	printLine("   c_c__/-c____/\n")
}

func PrintSimple(content string) {
	TextBlock(content)
	printLine("  \\\n")
	printLine("^..^      /\n")
	printLine("/_/\\_____/\n")
	printLine("   /\\   /\\\n")
	printLine("  /  \\ /  \\\n")
}

func PrintTazy(content string) {
	TextBlock(content)
	printLine("  \\                __\n")
	printLine("   \\___________   /  \\\n")
	printLine("               \\ / ..|\\\n")
	printLine("                (_\\  |_)\n")
	printLine("                /  \\@'\n")
	printLine("               /     \\\n")
	printLine("           _  /  `   |\n")
	printLine("          \\\\ /  \\  | _\\\n")
	printLine("           \\   /_ || \\\\_\n")
	printLine("            \\____)|_) \\_)\n")
}

func main() {
	args := os.Args

	if len(args) == 1 || len(args) > 4 {
		printLine("usage: dogsay [-b cur|maltipoo|simple|tazy] text\n")
		return
	}
	b := false
	choise := -1
	for i := 1; i < len(args); i++ {
		if args[i] == "-b" {
			b = true
			continue
		}
		if b && choise == -1 {
			if args[i] == "simple" {
				choise = 0
			} else if args[i] == "maltipoo" {
				choise = 1
			} else if args[i] == "tazy" {
				choise = 2
			} else if args[i] == "cur" {
				choise = 3
			} else {
				printLine("usage: dogsay [-b cur|maltipoo|simple|tazy] text\n")
				return
			}
		}
	}
	if b {
		if choise == -1 {
			printLine("usage: dogsay [-b cur|maltipoo|simple|tazy] text\n")
			return
		}
	}
	if !b {
		if len(args) > 2 {
			printLine("usage: dogsay [-b cur|maltipoo|simple|tazy] text\n")
			return
		}
		choise = 0
	}
	content := args[len(args)-1]
	if choise == 3 {
		PrintCur(content)
	} else if choise == 1 {
		PrintMaltipoo(content)
	} else if choise == 0 {
		PrintSimple(content)
	} else if choise == 2 {
		PrintTazy(content)
	}
}
