package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"strings"
)

func GenerateToken(length int) (string, error) {
	codes := make([]byte, length)
	if _, err := rand.Read(codes); err != nil {
		return "", err
	}

	for i := 0; i < length; i++ {
		codes[i] = uint8(48 + (codes[i] % 10))
	}
	return string(codes), nil
}
func GenerateRandomBytes(length int) ([]byte, error) {
	codes := make([]byte, length)
	if _, err := rand.Read(codes); err != nil {
		return []byte{}, err
	}

	for i := 0; i < length; i++ {
		codes[i] = uint8(48 + (codes[i] % 10))
	}
	return codes, nil
}

func GenerateRandomString(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}
func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Decode(s string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func EncodeToBase64String(data any) (string, error) {
	json_byte, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	encoded_string := Encode([]byte(json_byte))
	return encoded_string, nil
}

func DecodeBase64String(encoded_string string, data any) error {
	decoding, err := Decode(encoded_string)
	if err != nil {
		return err
	}
	err = json.Unmarshal(decoding, data)
	if err != nil {
		return err
	}
	return nil
}

var secret_key string

var iv string

func SetAuthParams(secret string) {
	strList := strings.Split(secret, "")[:32]
	secret_key = strings.Join(strList, "")
	strList = strings.Split(secret, "")[:16]
	iv = strings.Join(strList, "")
}

func Encrypt(text string, ivs ...string) (string, error) {
	enc_iv := iv
	if len(ivs) >= 1 {
		enc_iv = ivs[0]
	}
	block, err := aes.NewCipher([]byte(secret_key))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, []byte(enc_iv))
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)

	return Encode(cipherText), nil
}

// Decrypt method is to extract back the encrypted text
func Decrypt(text string, ivs ...string) (string, error) {
	dec_iv := iv
	if len(ivs) >= 1 {
		dec_iv = ivs[0]
	}
	block, err := aes.NewCipher([]byte(secret_key))
	if err != nil {
		return "", err
	}
	cipherText, err := Decode(text)
	if err != nil {
		return "", err
	}
	cfb := cipher.NewCFBDecrypter(block, []byte(dec_iv))
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
