package cipher

import (
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"io"
	"os"
	"testing"
)

func NewCipherBlockFuncN(key string) (cipher.Block, error) {
	return nil, errors.New("Error")
}

func ReadfullN(r io.Reader, buf []byte) (n int, err error) {
	return -1, errors.New("Error")
}

func DecodeStringN(s string) ([]byte, error) {
	return nil, errors.New("Error")
}

func encryptStreamFuncN(key string, iv []byte) (cipher.Stream, error) {
	return nil, errors.New("Error")
}

func decryptStreamFuncN(key string, iv []byte) (cipher.Stream, error) {
	return nil, errors.New("Error")
}
func TestEncryptStream(t *testing.T) {
	NewCipherBlockFunc = NewCipherBlockFuncN
	_, err := encryptStream("", nil)
	if err == nil {
		t.Error("Error Expected")
	}
	NewCipherBlockFunc = newCipherBlock
}

func TestEncrypt(t *testing.T) {
	key := "abc123"
	plaintext := "hello this is plaintext"
	_, err := Encrypt(key, plaintext)
	if err != nil {
		t.Error("Something went wrong!")
	}
	ReadFull = ReadfullN
	_, err = Encrypt("", plaintext)
	if err == nil {
		t.Error("Error Expected!")
	}
	ReadFull = io.ReadFull

	encryptStreamFunc = encryptStreamFuncN
	_, err = Encrypt("", plaintext)
	if err != nil {
		t.Error("Error Expected!")
	}
	encryptStreamFunc = encryptStream
}

func TestEncryptWriter(t *testing.T) {
	key := "123abc"
	w, err := os.OpenFile("/home/gs-1547/.secrets", os.O_RDWR|os.O_CREATE, 0755)
	_, err1 := EncryptWriter(key, w)
	if err != nil {
		t.Error("Something went Wrong!")
	}
	if err1 != nil {
		t.Error("Something went wrong!")
	}
	// Length of iv less than or null
	w, err = os.Open("/home/gs-1547/.secrets")
	_, err1 = EncryptWriter(key, w)
	if err != nil {
		t.Error("Something went Wrong!")
	}
	if err1 == nil {
		t.Error("Error Expected!")
	}

	// io.Readfull Error case
	ReadFull = ReadfullN
	_, err = EncryptWriter("", w)
	if err == nil {
		t.Error("Error Expected")
	}
	ReadFull = io.ReadFull

	// EncryptStream Error Case
	encryptStreamFunc = encryptStreamFuncN
	_, err = EncryptWriter(key, w)
	if err == nil {
		t.Error("Error Expected")
	}
	encryptStreamFunc = encryptStream
}

func TestDecryptStream(t *testing.T) {
	NewCipherBlockFunc = NewCipherBlockFuncN
	_, err := decryptStream("", nil)
	if err == nil {
		t.Error("Error Expected")
	}
	NewCipherBlockFunc = newCipherBlock
}

func TestDecrypt(t *testing.T) {
	key := "abc123"
	cipherHex := "acba5c457c8a2635cf9ce9bee2041b37c0d23d009005c0d178899c9e286309f784baa00cc732456015413bb34da44cfcc53348bef26cd6642bb7fa6c1be8417e109fb07b7abb1b98e671ef43653e21716f169b4ee90af4911c3116e2f734cc3a"
	_, err := Decrypt(key, cipherHex)
	if err != nil {
		t.Error("Something Went wrong!")
	}

	_, err = Decrypt(key, "")
	if err == nil {
		t.Error("Error Expected")
	}

	DecodeString = DecodeStringN
	_, err = Decrypt(key, cipherHex)
	if err == nil {
		t.Error("Error")
	}
	DecodeString = hex.DecodeString

	// decryptStream func error
	decryptStreamFunc = decryptStreamFuncN
	_, err = Decrypt(key, cipherHex)
	if err == nil {
		t.Error("Error")
	}
	decryptStreamFunc = decryptStream

}

func TestDecryptReader(t *testing.T) {
	// len(iv) null or less than expected!
	key := "123abc"
	r, err := os.Open("/home/gs-1547/test.txt")
	_, err1 := DecryptReader(key, r)
	if err != nil {
		t.Error("Something went Wrong !")
	}
	if err1 == nil {
		t.Error("Error Expected")
	}

	// Return statement
	r, err = os.Open("/home/gs-1547/old_perms.txt")
	_, err1 = DecryptReader(key, r)
	if err != nil {
		t.Error("Something went Wrong!")
	}
	if err1 != nil {
		t.Error("Something went wrong!")
	}

	// decryptStream func error
	decryptStreamFunc = decryptStreamFuncN

	_, err = DecryptReader(key, r)
	if err == nil {
		t.Error("Error Expected")
	}

	decryptStreamFunc = decryptStream
}
