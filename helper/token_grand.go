package helper

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func DoHashUsingSalt(text string) string {
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	var saltedText = fmt.Sprintf("text: '%s', salt: %s", text, salt)
	fmt.Println(saltedText)
	var sha = sha1.New()
	sha.Write([]byte(saltedText))
	var encrypted = sha.Sum(nil)
	return fmt.Sprintf("%x", encrypted)
}
