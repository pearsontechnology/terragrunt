package kmsgrunt

import (
	"fmt"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {

	svc, err := CreateKmsClient()
	if err != nil {
		t.Fatal(err)
	}
	decryptVal := Decrypt(svc, "ENCAQICAHjReWvTnkjTnSod2RSJwSHUFcyEWy0gM6aYJdklzXNllgGxzJVc7iFU+4gCkGZ7rZd7AAAAgzCBgAYJKoZIhvcNAQcGoHMwcQIBADBsBgkqhkiG9w0BBwEwHgYJYIZIAWUDBAEuMBEEDN/lBj6wQXm2Cn+W9wIBEIA/3IdagDE8ApOPMAjVK+Ff+HA99AzWB1rOLh4I9ecbhvb+t3K9rrcDw6vo+GP9e5Dir+v1q9+nPb7P6m4j2O1y")
	fmt.Println("decrypted val:", decryptVal)
	if decryptVal != "504b7ce4-cf52-4624-92b9-a54f6e32a4ce" {
		t.Fail()
	}
}
