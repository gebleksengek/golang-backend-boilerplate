package helpers

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"os"

	"../structs"
)

//MD5Hash generate md5 hash
func MD5Hash(text string) string {
	md5 := md5.New()
	md5.Write([]byte(text))
	return hex.EncodeToString(md5.Sum(nil))
}

//Base64Encoder generate base64 encoder
func Base64Encoder(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

//GeneratePassword hashing passwor
func GeneratePassword(data structs.Users) string {
	salt := os.Getenv("PASSWORD_SALT_KEY")
	salt = MD5Hash(salt)
	password := salt + data.Password
	password = MD5Hash(Base64Encoder(MD5Hash(Base64Encoder(password))))
	return password
}
