package encrypt

import (
	"bytes"
	"crypto/sha512"

	"golang.org/x/crypto/bcrypt"
)

const PasswordKey = "tQGUzh14eJRGRm7zO"

func Password(rawPassword string) (string, error) {
	hashPassword := password(rawPassword)
	pw, err := bcrypt.GenerateFromPassword(hashPassword, bcrypt.DefaultCost)
	return string(pw), err
}

func CompareHashAndPassword(hashedPassword string, rawPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), password(rawPassword))
}

func password(rawPassword string) []byte {
	hashPassword := sha512.Sum512([]byte(rawPassword))

	buf := bytes.NewBuffer(nil)
	buf.Write(hashPassword[:])
	buf.WriteString(PasswordKey)
	result := sha512.Sum512(buf.Bytes())

	return result[:]
}
