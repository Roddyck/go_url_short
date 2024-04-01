package encoder

import (
	"math/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const base = int64(len(charset))
const keyLenght = 6

func Encode() string {
    encoded := make([]byte, keyLenght)
    for i := range encoded {
        encoded[i] = charset[rand.Intn(len(charset))]
    }

    return string(encoded)
}
