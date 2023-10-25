package main

import (
	_ "embed"
	"fmt"
)

func DoWithInputPart01(cardPublicKey, doorPublicKey int) int {
	cardLoopSize := crackLoopSize(cardPublicKey)
	fmt.Printf("Cracked card loop size: %d\n", cardLoopSize)

	encryptionKey := transform(doorPublicKey, cardLoopSize)

	return encryptionKey
}

func crackLoopSize(publicKey int) int {
	v := 1
	loopSize := 0

	for v != publicKey {
		v = (v * 7) % 20201227
		loopSize++
	}

	return loopSize
}

func transform(subjectNumber, loopSize int) int {
	v := 1

	for i := 0; i < loopSize; i++ {
		v = (v * subjectNumber) % 20201227
	}

	return v
}

func DoWithInputPart02() int {
	return 0
}
