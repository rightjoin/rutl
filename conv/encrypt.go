package conv

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

// AES encryption
// http://codereview.stackexchange.com/questions/125846/encrypting-strings-in-golang

func DecryptBytes(data, key []byte) ([]byte, error) {
	// split the input up in to the IV seed and then the actual encrypted data.
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(data, data)
	return data, nil
}

// If consistent is true, then same key and input will always product the same output.
// If false, same key and input will produce a different encrypted output everytime.
func EncryptBytes(data, key []byte, consistent bool) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// create two 'windows' in to the output slice.
	output := make([]byte, aes.BlockSize+len(data))
	iv := output[:aes.BlockSize]
	encrypted := output[aes.BlockSize:]

	if !consistent {
		// populate the IV slice with random data.
		if _, err = io.ReadFull(rand.Reader, iv); err != nil {
			return nil, err
		}
	}

	stream := cipher.NewCFBEncrypter(block, iv)

	// note that encrypted is still a window in to the output slice
	stream.XORKeyStream(encrypted, data)
	return output, nil
}

// Decrypt takes two strings, cryptoText and keyString.
// cryptoText is the text to be decrypted and the keyString is the key to use for the decryption.
// The function will output the resulting plain text string with an error variable.
func Decrypt(cryptedText string, keyString string) (plainText string, err error) {

	encrypted, err := base64.URLEncoding.DecodeString(cryptedText)
	if err != nil {
		return "", err
	}
	if len(encrypted) < aes.BlockSize {
		return "", fmt.Errorf("cipherText too short. It decodes to %v bytes but the minimum length is 16", len(encrypted))
	}

	decrypted, err := DecryptBytes(encrypted, hashTo32Bytes(keyString))
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}

// Encrypt takes two string, plainText and keyString.
// plainText is the text that needs to be encrypted by keyString.
// The function will output the resulting crypto text and an error variable.
// If consistent is true, then same key and input will always produce the same output.
// If false, same key and input will produce a different encrypted output everytime.
func Encrypt(plainText string, keyString string, consistent bool) (cryptedText string, err error) {

	encrypted, err := EncryptBytes([]byte(plainText), hashTo32Bytes(keyString), consistent)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(encrypted), nil
}

func hashTo32Bytes(input string) []byte {

	data := sha256.Sum256([]byte(input))
	return data[0:]

}
