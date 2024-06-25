package firststruct

func (u User) PasswordReliability() string {
	if u.password == "" {
		return "undefined"
	}

	count := 0

	if len(u.password) >= 8 {
		count++
	}

	for _, up := range u.password {
		if 65 <= up && up <= 90 {
			count++
			break
		}
	}

	for _, low := range u.password {
		if 97 <= low && low <= 122 {
			count++
			break
		}
	}

	for _, digit := range u.password {
		if 48 <= digit && digit <= 57 {
			count++
			break
		}
	}

	for _, spec := range u.password {
		if !(48 <= spec && spec <= 57) && !(97 <= spec && spec <= 122) && !(65 <= spec && spec <= 90) {
			count++
			break
		}
	}

	switch count {
	case 5:
		return "strong"
	case 4, 3:
		return "medium"
	case 2, 1:
		return "easy"
	}

	return "undefined"
}
