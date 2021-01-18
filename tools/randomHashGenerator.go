package tools

import (
	"crypto/rand"
	"fmt"
)

//RandomHash Random Hash
func RandomHash() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
