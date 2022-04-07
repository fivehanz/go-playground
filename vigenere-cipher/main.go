/*
* Exercise from the book "Get Programming with Go"
* Unit 2
* Ref: https://en.wikipedia.org/wiki/Vigen%C3%A8re_cipher#Algebraic_description
* Implemented in UTF8 instead of ASCII.
 */

package main

import (
	"fmt"
	"math"
	"os"
	"unicode/utf8"
)

/*
* UTF8 def: https://www.rfc-editor.org/rfc/rfc3629#section-3
* utf8 can represent 2^21 chars
 */
var NoOfCharsInUTF8 int = int(math.Pow(2, 21))

/*
* []rune to string
* @return string
 */
func stringFromRuneArray(str []rune) string {
	//out := make([]rune, len(str))

	//for i := 0; i < len(str); i++ {
	//	out[i] = str[i]
	//}

	return string(str)
}

/*
* string to []rune
* @return []rune
 */
func runeArrayFromString(str string) []rune {
	//out := make([]rune, len(str))

	//for i := 0; i < len(str); i++ {
	//	out[i] = rune(str[i])
	//}

	return []rune(str)
}

/*
* E(Mi) = (Mi + K(i mod m)) mod l
* M -> Message, K -> Key, m -> len(K), l -> len(A)
 */
func encryptRuneArray(message []rune, key []rune) []rune {
	encrypted := make([]rune, len(message))

	for i := 0; i < len(message); i++ {
		encrypted[i] = rune(int(message[i]) + int(key[i%len(key)])%NoOfCharsInUTF8)
	}

	return (encrypted)
}

/*
* D(Ci) = (Ci - K(i mod m)) mod l
* C -> Cipher, K -> Key, m -> len(K), l -> len(A)
 */
func decryptRuneArray(cipher []rune, key []rune) []rune {
	decrypted := make([]rune, len(cipher))

	for i := 0; i < len(cipher); i++ {
		decrypted[i] = rune(int(cipher[i]) - int(key[i%len(key)])%NoOfCharsInUTF8)
	}

	return (decrypted)
}

// encrypt
func encrypt(message string, key string) string {
	// split str into runes
	runeMsg := runeArrayFromString(message)
	runeKey := runeArrayFromString(key)

	encRune := encryptRuneArray(runeMsg, runeKey)

	// join back as string
	return stringFromRuneArray(encRune)
}

// decrypt
func decrypt(message string, key string) string {
	// split into runes
	runeCipher := runeArrayFromString(message)
	runeKey := runeArrayFromString(key)

	decRune := decryptRuneArray(runeCipher, runeKey)

	// rune to str
	return stringFromRuneArray(decRune)
}

// for debugging
func printInASCII(text string) {
	for i := 0; i < len(text); i++ {
		fmt.Printf("%v -> %d -> %b \n", string(text[i]), text[i], text[i])
	}
}

// for debugging
func printInUTF8(text string) {
	for len(text) > 0 {
		//for i := 0; i < len(text); i++ {
		char, size := utf8.DecodeRune([]byte(text))
		fmt.Printf("%c -> %d -> %08b -> size: %v byte(s). \n", char, char, char, size)
		text = text[size:]
	}
}

func main() {
	sampleKey := "Lorem Isum ßåßˆøœ∆˚å∫ß∂˙∫∫ßåç≈"
	sampleText := "Eg: Go says so, that! A@B#C$- aå 水目馬 ;). πø∑ß∂œ∫√∆¥"

	if len(os.Args) > 1 {
		key := os.Args[1]

		if len(os.Args) == 2 {
			fmt.Printf("Key: %v\nSample-Text: %v\n\n", key, sampleText)

			text_enc := encrypt(sampleText, key)
			fmt.Printf("Encrypted Text: %v\n", text_enc)

			text_dec := decrypt(text_enc, key)
			fmt.Printf("Decrypted Text: %v\n", text_dec)
		} else {
			str := os.Args[2]
			fmt.Printf("Key: %v\nText: %v\n\n", key, str)

			text_enc := encrypt(str, key)
			fmt.Printf("Encrypted Text: %v\n", text_enc)

			text_dec := decrypt(text_enc, key)
			fmt.Printf("Decrypted Text: %v\n", text_dec)
		}

	} else {
		fmt.Println("usage:\n$ vigenere-cipher KEY MESSAGE\nsample test case below:\n")

		fmt.Printf("Sample-Key: %v\nSample-Text: %v\n\n", sampleKey, sampleText)

		text_enc := encrypt(sampleText, sampleKey)
		fmt.Printf("Encrypted Text: %v\n", text_enc)

		text_dec := decrypt(text_enc, sampleKey)
		fmt.Printf("Decrypted Text: %v\n", text_dec)
	}

}
