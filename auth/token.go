package auth

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
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
