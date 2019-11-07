package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
)

var NewCipherBlockFunc = newCipherBlock
var ReadFull = io.ReadFull
var DecodeString = hex.DecodeString
var decryptStreamFunc = decryptStream
var encryptStreamFunc = encryptStream

// This code is based the standard library example at -
// https://golang.org/pkg/crypto/cipher/#NewCFBEncrypter
func encryptStream(key string, iv []byte) (cipher.Stream, error) {
	block, err := NewCipherBlockFunc(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewCFBEncrypter(block, iv), nil

}

// Encrypt takes in a key and plaintext and return a hex representation of the encrypted value.
// The reason why we are taking key as string,
// so that the user does not need to think of the length of the string and how can they convert it to suitable size.
func Encrypt(key, plaintext string) (string, error) {
	// The Initialization Vector needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream, err := encryptStreamFunc(key, iv)
	if err != nil {
		return "", nil
	}
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	// Turning it into a Hex Value.
	return fmt.Sprintf("%x", ciphertext), nil
}

// EncryptWriter will return a writer that will write encrypted data to the original writer.
func EncryptWriter(key string, w io.Writer) (*cipher.StreamWriter, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream, err := encryptStreamFunc(key, iv)
	if err != nil {
		return nil, err
	}
	n, err := w.Write(iv)
	if n != len(iv) || err != nil {
		log.Printf("error :%v len :%v", err, len(iv))
		return nil, errors.New("Encrypt:Unable to write full iv to writer")
	}
	return &cipher.StreamWriter{S: stream, W: w}, nil
}
func decryptStream(key string, iv []byte) (cipher.Stream, error) {
	block, err := NewCipherBlockFunc(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewCFBDecrypter(block, iv), nil
}

// Decrypt - Takes in a key and a cipherHex(hex representation of
// the ciphertext) and decrypt it.
// This code is based on the standard library example at -
// https://golang.org/pkg/crypto/cipher/#NewCFBDecrypter
func Decrypt(key, cipherHex string) (string, error) {
	//cipherHex is the Encrypted value
	//DecodeString converts that hex value into a byte slice.
	ciphertext, err := DecodeString(cipherHex)
	if err != nil {
		return "", err
	}
	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("Encrypt: Cipher too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream, err := decryptStreamFunc(key, iv)
	if err != nil {
		return "", err
	}
	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil
}

//DecryptReader will return a reader that will decrypt data from the provided reader
// and give the user a way to read the data as if it was not encrypted
func DecryptReader(key string, r io.Reader) (*cipher.StreamReader, error) {
	iv := make([]byte, aes.BlockSize)
	n, err := r.Read(iv)
	if n < len(iv) || err != nil {
		return nil, errors.New("Encrypt: Unable to read the full iv")
	}
	stream, err := decryptStreamFunc(key, iv)
	if err != nil {
		return nil, err
	}
	return &cipher.StreamReader{S: stream, R: r}, nil
}

func newCipherBlock(key string) (cipher.Block, error) {
	hasher := md5.New()
	fmt.Fprint(hasher, key)
	cipherkey := hasher.Sum(nil)
	return aes.NewCipher(cipherkey)
}
