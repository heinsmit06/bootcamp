package bootcamp

//
// func Atoi(s string) int {
// 	sum := 0
//
// 	if (s[0] == '-') || (s[0] == '+') {
//
// 		for i := 1; i < len(s); i++ {
// 			not_valid_flag := 0
// 			if !(48 <= rune(s[i]) && rune(s[i]) <= 57) {
// 				not_valid_flag++
// 			}
//
// 			if not_valid_flag == 1 {
// 				return 0
// 			}
// 		}
//
// 		var max_degree_of_int float64 = float64(len(s) - 2)
//
// 		for i := 1; i < len(s); i++ {
// 			sum += (int(s[i]) - 48) * Pow(10, max_degree_of_int)
// 			max_degree_of_int--
// 		}
//
// 		if s[0] == '-' {
// 			return sum * (-1)
// 		} else {
// 			return sum
// 		}
//
// 	}
//
// 	for i := 0; i < len(s); i++ {
// 		not_valid_flag := 0
// 		if !(48 <= rune(s[i]) && rune(s[i]) <= 57) {
// 			not_valid_flag++
// 		}
// 		if not_valid_flag == 1 {
// 			return 0
// 		}
// 	}
//
// 	var max_degree_of_int float64 = float64(len(s) - 1)
//
// 	for i := 0; i < len(s); i++ {
// 		sum += (int(s[i]) - 48) * Power(10, max_degree_of_int)
// 		max_degree_of_int--
// 	}
//
// 	return sum
// }

//func main() {
//	fmt.Println(Atoi("123"))
//	fmt.Println(Atoi("+123"))
//	fmt.Println(Atoi("-123"))
//	fmt.Println(Atoi("-123!"))
//	fmt.Println(Atoi("abc"))
//}

//func Power(n, degree int) int {
//	power := 0
//	for i := 0; i < degree; i++ {
//		power *=
//	}
//}
