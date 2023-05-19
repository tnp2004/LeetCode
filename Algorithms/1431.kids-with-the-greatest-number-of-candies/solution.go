package kidswiththegreatestnumberofcandies

func kidsWithCandies(candies []int, extraCandies int) []bool {

	var greatestKid = make([]bool, len(candies))

	maximumCandie := fineMaximum(candies)

	for i, candie := range candies {
		if candie+extraCandies >= maximumCandie {
			greatestKid[i] = true
		}
	}

	return greatestKid

}

func fineMaximum(s []int) int {
	maxVal := 0
	for _, v := range s {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}
