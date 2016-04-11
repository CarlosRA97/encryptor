package encrypt

import (
	"crypto/sha1"
	"fmt"
)

// HashIt This funcion hash the string given
// and send it directly to the clipboard
func HashIt(text string) string {
	const question = "Do you want to copy to clipboard?"
	data := []byte(text)
	hash := fmt.Sprintf("%x\n", sha1.Sum(data))
	ClipboardIt(hash)
	return hash
}
