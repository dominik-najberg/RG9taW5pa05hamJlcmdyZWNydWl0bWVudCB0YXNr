package internal

import (
	"encoding/base64"
)

func CalculateBase64() string {
	msg := "Dominik" + "Najberg" + "recruitment task"
	return base64.StdEncoding.EncodeToString([]byte(msg))
}
