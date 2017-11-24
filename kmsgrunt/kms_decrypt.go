package kmsgrunt

import (
	"encoding/base64"
	"fmt"

	"github.com/aws/aws-sdk-go/service/kms"
)

// Decrypt decrypts a passed argument
func Decrypt(svc *kms.KMS, encryptedVal string) string {
	blob, _ := base64.StdEncoding.DecodeString(encryptedVal[3:])

	// Decrypt the data
	result, err := svc.Decrypt(&kms.DecryptInput{CiphertextBlob: blob})

	if err != nil {
		fmt.Println("Error decrypting data: ", err)
		return encryptedVal
	}

	blobString := string(result.Plaintext)

	return (blobString)
}
