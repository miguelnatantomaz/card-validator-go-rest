package validator

func ValidateCard(number string) bool {
	var sum int
	var alternate = false

	for i := len(number) - 1; i>= 0; i-- {
		digit := int(number[i] - '0')

		if alternate {
			digit *= 2
			if digit > 9 {
				// var sumDigit int = 0
				// for _, char := range strconv.Itoa(digit) {
				// 	sumDigit += int(char - '0')
				// }
				// digit = sumDigit

				digit = (digit % 10) + (digit / 10)
			}
		}

		alternate = !alternate
		sum += digit
	}
	
	return sum%10 == 0
}