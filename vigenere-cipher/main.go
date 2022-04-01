/*
* Exercise from the book "Get Programming with Go"
* Unit 2
* Ref: https://en.wikipedia.org/wiki/Vigen%C3%A8re_cipher#Algebraic_description
 */

package main

import "fmt"

/*
* []byte to string
* @return string
 */
func stringFromByteArray(str []byte) string {
	out := make([]rune, len(str))

	for i := 0; i < len(str); i++ {
		out[i] = rune(str[i])
	}

	return string(out)
}

/*
* string to []byte
* @return []byte
 */
func byteArrayFromString(str string) []byte {
	out := make([]byte, len(str))

	for i := 0; i < len(str); i++ {
		out[i] = byte(str[i])
	}

	return out
}

/*
* E(Mi) = (Mi + K(i mod m)) mod l
* M -> Message, K -> Key, m -> len(K), l -> len(A)
 */
func encrypt(message string, key string) []byte {
	encrypted := make([]byte, len(message))

	for i := 0; i < len(message); i++ {
		encrypted[i] = byte((rune(message[i]) + rune(key[i%len(key)])))
	}

	return (encrypted)
}

/*
* D(Ci) = (Ci - K(i mod m)) mod l
* C -> Cipher, K -> Key, m -> len(K), l -> len(A)
 */
func decrypt(cipher string, key string) []byte {
	decrypted := make([]byte, len(cipher))

	for i := 0; i < len(cipher); i++ {
		decrypted[i] = byte((rune(cipher[i]) - rune(key[i%len(key)])))
	}

	return (decrypted)
}

func main() {
	key := "Jonz-Is AName"
	text := "Hanz Is A@B#C$- Chocolate."
	fmt.Println(text, byteArrayFromString(text), stringFromByteArray(byteArrayFromString(text)))
	text_enc := encrypt(text, key)
	fmt.Println(stringFromByteArray(text_enc), text_enc)
	text_dec := decrypt(stringFromByteArray(text_enc), key)
	fmt.Println(stringFromByteArray(text_dec), text_dec)
}
