// Placeholder
package main

func main() {
	// Placeholder
}

// Shift letter
//
// Shift a letter right by the given number.
// Wrap the letter around if it reaches the end of the alphabet.
//
// Examples
// shiftLetter("A", 0) -> "A"
// shiftLetter("A", 2) -> "C"
// shiftLetter("Z", 1) -> "A"
// shiftLetter("X", 5) -> "C"
// shiftLetter(" ", _) -> " "
//
// Note: we use the underscore `_` to denote the presence of a value that is present but irrelevant.
//
// Params:
// - letter, a single uppercase English letter, or a space
// - shift, the number by which to shift the letter
//
// Returns:
// - the letter, shifted appropriately, if a letter; otherwise a space.
func shiftLetter(letter string, shift int) string {
	if letter == " " {
		return " "
	}

	var outputLetter int = ((int(letter[0]) - 'A' + shift) % 26) + 'A'

	return string(outputLetter)
}

// Caesar cipher
//
// Apply a shift number to a string of uppercase English letters and spaces.
//
// Params:
// - message, a string of uppercase English letters and spaces
// - shift, the number by which to shift the letters
//
// Returns
// - the message, shifted appropriately
func caesarCipher(message string, shift int) string {
	var final string = ""

	for _, letter := range message {
		if string(letter) == " " {
			final += " "
		} else {
			var outputLetter int = ((int(letter) - 'A' + shift) % 26) + 'A'
			final += string(outputLetter)
		}
	}
	return final
}

// Shift by letter
//
// Shift a letter to the right using the number equivalent of another letter.
// The shift letter is any letter from A to Z, where A represents 0, B represents 1, ..., Z represents 25.
//
// Examples
// shiftByLetter("A", "A") -> "A"
// shiftByLetter("A", "C") -> "C"
// shiftByLetter("B", "K") -> "L"
// shiftByLetter(" ", _) -> " "
//
// Params:
// - letter, a single uppercase English letter, or a space
// - letterShift, a single uppercase English letter
//
// Returns:
// - the letter, shifted appropriately
func shiftByLetter(letter string, letterShift string) string {
	if letter == " " {
		return " "
	}

	var shift int = int(letterShift[0]) - 'A'
	var outputLetter int = ((int(letter[0]) - 'A' + shift) % 26) + 'A'

	return string(outputLetter)
}

// Vigenere cipher
//
// Encrypt a message using a keyphrase instead of a static number.
// Every letter in the message is shifted by the number represented by the respective letter in the key.
// Spaces are ignored.
//
// Example
// vigenereCipher("A C", "KEY") -> "K A"
//
// If needed, the keyphrase is extended to match the length of the key.
// If the key is "KEY" and the message is "LONGTEXT", the key will be extended to "KEYKEYKE".
//
// Params:
// - message, a string of uppercase english letters and/or spaces
// - key, a string of uppercase English letters with no spaces. Will not exceed the length of the message.
//
// Returns:
// - the message, shifted appropriately
func vigenereCipher(message string, key string) string {
	var final string = ""

	var keyIndex int = 0
	for _, letter := range message {

		if letter == ' ' {
			final += " "
		} else {
			var shift int = int(key[keyIndex]) - 'A'
			var outputLetter int = ((int(letter) - 'A' + shift) % 26) + 'A'
			final += string(outputLetter)
		}
		keyIndex = (keyIndex + 1) % len(key)
	}
	return final
}

// Scytale cipher
//
// Encrypts a message by simulating a scytale cipher.
//
// A scytale is a cylinder around which you can wrap a long strip of
//
//	parchment that contained a string of apparent gibberish. The parchment,
//	when read using the scytale, would reveal a message due to every nth
//	letter now appearing next to each other, revealing a proper message.
//
// This encryption method is obsolete and should never be used to encrypt
//
//	data in production settings.
//
// You may read more about the method here:
//
//	https://en.wikipedia.org/wiki/Scytale
//
// You may follow this algorithm to implement a scytale-style cipher:
//  1. Take a message to be encoded and a "shift" number.
//     For this example, we will use "INFORMATION_AGE" as
//     the message and 3 as the shift.
//  2. Check if the length of the message is a multiple of
//     the shift. If it is not, add additional underscores
//     to the end of the message until it is.
//     In this example, "INFORMATION_AGE" is already a multiple of 3,
//     so we will leave it alone.
//  3. This is the tricky part. Construct the encoded message.
//     For each index i in the encoded message, use the character at the index
//     (i // shift) + (len(message) // shift) * (i % shift) of the raw message.
//     If this number doesn't make sense, you can play around with the cipher at
//     https://dencode.com/en/cipher/scytale to try to understand it.
//  4. Return the encoded message. In this case,
//     the output should be "IMNNA_FTAOIGROE".
//
// Example
// scytaleCipher("INFORMATION_AGE", 3) -> "IMNNA_FTAOIGROE"
// scytaleCipher("INFORMATION_AGE", 4) -> "IRIANMOGFANEOT__"
// scytaleCipher("ALGORITHMS_ARE_IMPORTANT", 8) -> "AOTSRIOALRH_EMRNGIMA_PTT"
//
// Params:
// - message, a string of uppercase English letters and underscores. Underscores represent spaces.
// - shift, a positive integer that does not exceed the length of the message.
//
// Returns:
// - the message, encoded appropriately.
func scytaleCipher(message string, shift int) string {
	var currentMessage string = message

	for len(currentMessage)%shift != 0 {
		currentMessage += "_"
	}
	rows := len(currentMessage) / shift

	result := make([]byte, len(currentMessage))
	for i := 0; i < len(currentMessage); i++ {
		originalIndex := (i / shift) + rows*(i%shift)
		result[i] = currentMessage[originalIndex]
	}

	return string(result)
}

// Scytale decipher
//
// Decrypts a message that was originally encrypted with the `scytaleCipher` function above.
//
// Example:
// scytaleDecipher('IMNNA_FTAOIGROE', 3) -> 'INFORMATION_AGE'
// scytaleDecipher('AOTSRIOALRH_EMRNGIMA_PTT', 8) -> 'ALGORITHMS_ARE_IMPORTANT'
// scytaleDecipher('IRIANMOGFANEOT__', 4) -> 'INFORMATION_AGE_'
//
// Params:
// - message, a string of uppercase English letters and underscores. Underscores represent spaces.
// - shift, a positive integer that does not exceed the length of the message.
//
// Returns:
// - The message, decoded appropriately.
func scytaleDecipher(message string, shift int) string {
	var legth int = len(message)
	result := make([]rune, legth)
	for i, letter := range message {
		result[(i%shift)*legth/shift+i/shift] = letter
	}
	return string(result)
}
