package helper

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/spf13/viper"
)

func CreateSign(viper *viper.Viper, method string, requestPath string, body string, timestamp string) string {
	secret := []byte(viper.GetString("secrets.OKX_SECRET_KEY"))
	data := []byte(timestamp + method + requestPath + body)
	hmacNew := hmac.New(sha256.New, secret)
	hmacNew.Write(data)

	return hex.EncodeToString(hmacNew.Sum(nil))
}
