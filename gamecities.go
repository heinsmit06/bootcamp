package bootcamp

func GameCities(cities []string) []string {
	if len(cities) == 0 {
		return []string{}
	}
	var dfs func(currentCity string, remainingCities []string, currentChain []string) []string
	dfs = func(currentCity string, remainingCities []string, currentChain []string) []string {
		rawCurrentCity := []rune(currentCity)
		currentChain = append(currentChain, currentCity)

		if len(remainingCities) == 0 {
			return currentChain
		}

		longestChain := currentChain
		lastChar := LowerCaseLetter(rawCurrentCity[len(rawCurrentCity)-1])

		for i, nextCity := range remainingCities {
			rawNextCity := []rune(nextCity)

			if LowerCaseLetter(rawNextCity[0]) == lastChar {
				newRemainingCities := append([]string{}, remainingCities[:i]...)
				newRemainingCities = append(newRemainingCities, remainingCities[i+1:]...)
				newChain := dfs(nextCity, newRemainingCities, currentChain)
				if len(newChain) > len(longestChain) {
					longestChain = newChain
				}
			}
		}
		return longestChain
	}

	longestOverallChain := []string{}

	for i, city := range cities {
		remainingCities := append([]string{}, cities[:i]...)
		remainingCities = append(remainingCities, cities[i+1:]...)
		chain := dfs(city, remainingCities, []string{})
		if len(chain) > len(longestOverallChain) {
			longestOverallChain = chain
		}
	}

	return longestOverallChain
}

func LowerCaseLetter(character rune) rune {
	characterIndex := int(character)
	if characterIndex >= int('A') && characterIndex <= int('Z') {
		return rune(characterIndex + 32)
	}
	return character
}
