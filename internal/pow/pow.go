package pow

import (
	"crypto/sha1"
	"fmt"
	"io"
)

const zeroByte = 48

type Pow struct {
	Data  string
	Nonce int
}

func (h Pow) String() string {
	return fmt.Sprintf("%s:%d", h.Data, h.Nonce)
}

func IsHashCorrect(hash string, zerosCount int) bool {
	if zerosCount > len(hash) {
		return false
	}
	for _, ch := range hash[:zerosCount] {
		if ch != zeroByte {
			return false
		}
	}
	return true
}

func (h Pow) Hash(maxNonce int, zerosCount int) (hash string, err error) {
	for h.Nonce <= maxNonce {
		hash := makeHash(h.String())
		if IsHashCorrect(hash, zerosCount) {
			return hash, nil
		}
		h.Nonce++
	}
	return "", fmt.Errorf("cannot generate hash")
}

func makeHash(data string) string {
	h := sha1.New()
	_, _ = io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
