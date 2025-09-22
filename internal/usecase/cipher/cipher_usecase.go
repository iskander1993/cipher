package cipher

import (
	"ave_project/pkg/cipher"
)

type CipherUsecase struct{}

func (c *CipherUsecase) Encrypt(text string, shift int) string {
	return cipher.Caesar(text, shift)
}

func (c *CipherUsecase) Decrypt(text string, shift int) string {
	return cipher.Caesar(text, -shift)
}
