package main

import (
	_ "embed"
)

type World struct {
	From, To int
}

func DoWithInputPart01(world World) int {
	password := toBytes(world.From)

	count := 0
	for {
		if toNumber(password) > world.To {
			break
		}

		if validate(password) {
			count++
		}

		if !increment(password) {
			break
		}
	}

	return count
}

func increment(password []byte) bool {
	for i := len(password) - 1; i >= 0; i-- {
		if password[i] < 9 {
			password[i]++
			return true
		}

		if i == 0 {
			return false
		}

		password[i] = 0
	}

	return true
}

func toBytes(number int) []byte {
	return []byte{
		byte(number / 100000),
		byte(number / 10000 % 10),
		byte(number / 1000 % 10),
		byte(number / 100 % 10),
		byte(number / 10 % 10),
		byte(number % 10),
	}
}

func toNumber(password []byte) int {
	return int(password[0])*100000 +
		int(password[1])*10000 +
		int(password[2])*1000 +
		int(password[3])*100 +
		int(password[4])*10 +
		int(password[5])
}

func validate(password []byte) bool {
	lastDigit := password[0]
	sameDigits := false

	for i, digit := range password {
		// never decrease rule
		if digit < lastDigit {
			return false
		}

		// two adjacent digits are the same rule
		if i > 0 && digit == lastDigit {
			sameDigits = true
		}

		lastDigit = digit
	}

	return sameDigits
}

func DoWithInputPart02(world World) int {
	password := toBytes(world.From)

	count := 0
	for {
		if toNumber(password) > world.To {
			break
		}

		if validate2(password) {
			count++
		}

		if !increment(password) {
			break
		}
	}

	return count
}

func validate2(password []byte) bool {
	lastDigit := password[0]
	sameDigits := map[byte]byte{}

	for i, digit := range password {
		// never decrease rule
		if digit < lastDigit {
			return false
		}

		// two adjacent digits are the same rule
		if i > 0 && digit == lastDigit {
			sameDigits[digit]++
		}

		lastDigit = digit
	}

	for _, count := range sameDigits {
		if count == 1 {
			return true
		}
	}

	return false
}
